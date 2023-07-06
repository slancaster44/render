
#include "immintrin.h"


inline __m256d Add(const double* t1, const double* t2) {
    __m256d v1 = _mm256d_load_pd(t1);
    __m256d v2 = _m256d_load_pd(t2);
    return _mm256_add_pd(v1, v2);
}