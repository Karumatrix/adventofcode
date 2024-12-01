#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <fcntl.h>
#include <time.h>
#include <sys/time.h>

int getLineCount(FILE* fd);
int isDigit(int ch);
int sizeOfNumber(int number);
int sizeOfLong(long int number);
int stringEqual(char* str1, char* str2);

#define timerStart clock_t __start = clock(); \
    struct timeval __startTv; \
    gettimeofday(&__startTv, NULL);
#define timerEnd clock_t __end = clock(); \
    struct timeval __endTv; \
    gettimeofday(&__endTv, NULL); \
    clock_t __cpuTime = __end - __start; \
    long double __time = ((1000000 * __endTv.tv_sec) + __endTv.tv_usec) - ((1000000 * __startTv.tv_sec) + __startTv.tv_usec); \
    printf("CPU time: %ld, in seconds: %.6Lfs\n", __cpuTime, __time/1000000);
    
int isDigit(int ch){
    if(ch >= '0' && ch <= '9') {
        return 1;
    }
    return 0;
}

int sizeOfNumber(int number) {
    int i = 1;
    while(number >= 10){
        number = number / 10;
        i++;
    }
    return i;
}

int sizeOfLong(long int number){
    int i = 1;
    while(number >= 10){
        number = number / 10;
        i++;
    }
    return i;
}

int getLineCount(FILE* fd){
    int ch;
    int lineCount = 0;
    int position = ftell(fd);

    do {
        ch = fgetc(fd);
        if(ch == '\n'){
            lineCount++;
        }
    } while(ch != EOF);

    fseek(fd, position, SEEK_SET);
    return lineCount;
}

int stringEqual(char* str1, char* str2){
    int size = strlen(str1);
    if(size != strlen(str2)) return 0;
    for(int i = 0; i < size; i++){
        if(str1[i] != str2[i]) return 0;
    }
    return 1;
}

long int* seedList;
char** lines;
char* mapNames[] = {"seed-to-soil map:\n", "soil-to-fertilizer map:\n",
                "fertilizer-to-water map:\n", "water-to-light map:\n",
                "light-to-temperature map:\n", "temperature-to-humidity map:\n",
                "humidity-to-location map:\n"};
int mapBeginningLines[8];
int seedPairCount;
int lineCount;
FILE* fd;

long int handleSeed(int seedNumber);

int main(int argc, char *argv[]) {

    if(argc != 2) {
        printf("Usage: %s <input_file>\n", argv[0]);
        return -1;
    }

    fd = fopen(argv[1], "r");

    if (fd == NULL) {
        printf("File not found\n");
        return -1;
    }

    timerStart;

    lineCount = getLineCount(fd);

    seedList = malloc(sizeof(int)*50);
    lines = malloc(sizeof(char*)*lineCount-1);

    for(int i = 0; i < lineCount; i++){
        size_t size = 320;
        lines[i] = malloc(size);
        getline(&lines[i], &size, fd);
    }

    int seedLineSize = strlen(lines[0]);

    int seedLineIndex = strlen("seeds: ");
    seedPairCount = 0;

    /* Search for all seed numbers and store them in seedList */
    while(seedLineIndex < seedLineSize) {
        if(isDigit(lines[0][seedLineIndex]) == 1){
            long int number;
            sscanf((char*) (lines[0] + seedLineIndex), "%ld", &number);
            seedList[seedPairCount] = number; 
            seedPairCount++;
            seedLineIndex += sizeOfLong(number);
        } else {
        seedLineIndex++;
        }
    }

    for(int i = 0; i < lineCount; i++){
        char* line = lines[i];
        for(int j = 0; j < sizeof(mapNames)/sizeof(char*); j++){
            if(stringEqual(mapNames[j], line) == 1){
                mapBeginningLines[j] = i;
            }
        }
    }
    mapBeginningLines[7] = lineCount;

    long int lowestLocation = 0;
    for(int i = 0; i < seedPairCount; i++){
        long int location = handleSeed(i);
        if(location < lowestLocation || i == 0) lowestLocation = location;
    }

    printf("Solution: %ld\n", lowestLocation);

    timerEnd;
}

long int handleSeed(int seedNumber){
    long int seed = seedList[seedNumber];
    
    for(int i = 0; i < sizeof(mapBeginningLines)/sizeof(int); i++){
        for(int line = mapBeginningLines[i]+1; line < mapBeginningLines[i+1]; line++){
            long int destination = 0;
            long int source = 0;
            long int length = 0;
            sscanf(lines[line], "%ld %ld %ld", &destination, &source, &length);
            if(seed >= source && seed <= source+length){
                long int move = destination - source;
                seed = seed + move;
                break;
            }
        }
    }

    return seed;
}
