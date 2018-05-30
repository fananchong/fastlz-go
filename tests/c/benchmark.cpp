#include "Contrib/fastlz/fastlz.h"
#include <chrono>
#include <cassert>
#include <stdio.h>

long long get_tick_count(void)
{
    typedef std::chrono::time_point<std::chrono::system_clock, std::chrono::nanoseconds> nanoClock_type;
    nanoClock_type tp = std::chrono::time_point_cast<std::chrono::nanoseconds>(std::chrono::system_clock::now());
    return tp.time_since_epoch().count();
}

const int RAND_MAX_COUNT = 200000 * sizeof(float);
char randValue[RAND_MAX_COUNT];

const int count = 10000;

char out[count][4096 * 2];
int outLen[count];
char dec[count][4096 * 2];

int main() {
    FILE* f1 = fopen("../../rand.bin", "rb");
    fread(randValue, RAND_MAX_COUNT, 1, f1);
    fclose(f1);

    int dataLen[] = { 4096,128,512,1024,4096 };
    printf("total count: %d\n", count);

    for (size_t k = 0; k < sizeof(dataLen) / sizeof(dataLen[0]); k++)
    {
        printf("data size:           %20dbyte\n", dataLen[k]);

        auto t1 = get_tick_count();
        for (int i = 0; i < count; i++) {
            outLen[i] = fastlz_compress(randValue + i, dataLen[k], out[i]);
            assert(outLen[i] != 0);
        }
        auto t2 = get_tick_count();
        printf("compress cost:        %20lldns %20.3fns/op %20.3fms/op\n",
            t2 - t1, float(t2 - t1) / count, float(t2 - t1) / count / 1000000);

        t1 = get_tick_count();
        for (int i = 0; i < count; i++) {
            int ln = fastlz_decompress(out[i], outLen[i], dec[i], sizeof(dec) / sizeof(dec[0]));
            assert(ln != 0);
        }
        t2 = get_tick_count();
        printf("decompress cost:      %20lldns %20.3fns/op %20.3fms/op\n",
            t2 - t1, float(t2 - t1) / count, float(t2 - t1) / count / 1000000);
    }
}
