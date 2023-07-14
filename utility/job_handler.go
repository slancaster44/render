package utility

import (
	"reflect"
	"sync"
)

func DoJobs[T any](list []T, number_jobs int, function func(item T) T) []T {
	job_chunk_size := len(list) / number_jobs
	leftovers := len(list) % number_jobs

	ch := make(chan T)

	for i := 0; i < number_jobs; i++ {

		go func(i int) {

			chunk_start := i * job_chunk_size
			chunk_end := chunk_start + job_chunk_size

			for j := chunk_start; j < chunk_end; j++ {
				ch <- function(list[j])
			}

		}(i)
	}

	output := []T{}
	for i := 0; i < job_chunk_size*number_jobs; i++ {
		job_output := <-ch
		output = append(output, job_output)
	}

	leftovers_start := job_chunk_size * number_jobs
	for i := leftovers_start; i < leftovers_start+leftovers; i++ {
		output = append(output, function(list[i]))
	}

	return output
}

func CallOn[T any](list []T, number_jobs int, function func(item T)) {
	job_chunk_size := len(list) / number_jobs
	leftovers := len(list) % number_jobs

	var wg sync.WaitGroup

	for i := 0; i < number_jobs; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			chunk_start := i * job_chunk_size
			chunk_end := chunk_start + job_chunk_size

			for j := chunk_start; j < chunk_end; j++ {
				function(list[j])
			}

		}(i)
	}

	leftovers_start := job_chunk_size * number_jobs
	for i := leftovers_start; i < leftovers_start+leftovers; i++ {
		function(list[i])
	}

	wg.Wait()

}

func MethodOn[T any](list []T, number_jobs int, method_name string) {
	job_chunk_size := len(list) / number_jobs
	leftovers := len(list) % number_jobs

	var wg sync.WaitGroup

	for i := 0; i < number_jobs; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			chunk_start := i * job_chunk_size
			chunk_end := chunk_start + job_chunk_size

			for j := chunk_start; j < chunk_end; j++ {
				reflect.ValueOf(list[j]).MethodByName(method_name).Call([]reflect.Value{})
			}

		}(i)
	}

	leftovers_start := job_chunk_size * number_jobs
	for i := leftovers_start; i < leftovers_start+leftovers; i++ {
		reflect.ValueOf(list[i]).MethodByName(method_name).Call([]reflect.Value{})
	}

	wg.Wait()

}
