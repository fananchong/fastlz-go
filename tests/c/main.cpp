#include "Contrib/fastlz/fastlz.h"
#include <cassert>
#include <string>
#include <vector>
#include <time.h>

const int RAND_MAX_COUNT = 200000;
float randValue[RAND_MAX_COUNT];
int randIndex = 0;
inline float frand()
{
    auto v = randValue[randIndex];
    randIndex++;
    return v;
}

const int OUT_MAX_COUNT = 10000;
float outValue[OUT_MAX_COUNT];
int outIndex = 0;


float myrand() {
    return (float)rand() / (float)RAND_MAX;
}

const char* MESH_FILE = "../../nav_test.obj.tile.bin";

int main(int argn, char* argv[]) {
    srand((unsigned int)(time(0)));

    if (argn > 1 && argv[1] == std::string("rand")) {
        FILE* f = fopen("../../rand.bin", "wb");
        for (int i = 0; i < RAND_MAX_COUNT;i++) {
            float v = (float)rand() / (float)RAND_MAX;
            fwrite(&v, sizeof(float), 1, f);
        }
        fclose(f);
    }
    else {
        FILE* f1 = fopen("../../rand.bin", "rb");
        fread(randValue, RAND_MAX_COUNT * sizeof(float), 1, f1);
        fclose(f1);
        void test();
        test();
        FILE* f2 = fopen("../../result.bin", "wb");
        fwrite(outValue, outIndex * sizeof(float), 1, f2);
        fclose(f2);
    }
    return 0;
}

void test() {
    const int size = RAND_MAX_COUNT * sizeof(float)*1.06f;
    {
        char out[size];
        int len = fastlz_compress(randValue, RAND_MAX_COUNT * sizeof(float), out);
        int v = 0;
        for (size_t i = 0; i < len; i++) {
            v += out[i];
        }
        outValue[outIndex++] = float(v);
    }

    {
        char out[size];
        int len = fastlz_compress_level(1, randValue, RAND_MAX_COUNT * sizeof(float), out);
        int v = 0;
        for (size_t i = 0; i < len; i++) {
            v += out[i];
        }
        outValue[outIndex++] = float(v);
    }

    {
        char out[size];
        int len = fastlz_compress_level(2, randValue, RAND_MAX_COUNT * sizeof(float), out);
        int v = 0;
        for (size_t i = 0; i < len; i++) {
            v += out[i];
        }
        outValue[outIndex++] = float(v);
    }
}