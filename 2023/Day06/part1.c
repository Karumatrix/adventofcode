#include <fcntl.h>
#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>
#include <sys/stat.h>
#include <stddef.h>
#include <string.h>

char **load_2d_arr_from_file(char *filecontent)
{
    int num_line = 0;
    int num_char = 0;
    int index = 0;
    int temp_index = 0;
    int result_index = 0;
    int temp_result_index = 0;
    while (filecontent[index] != '\0') {
        if (filecontent[index] == '\n')
            num_line += 1;
        index += 1;
    }
    char **result = malloc(sizeof(char *) * (num_line + 2));
    index = 0;
    while (filecontent[index] != '\0') {
        temp_index = index;
        num_char = 0;
        while (filecontent[temp_index] != '\0') {
            if (filecontent[temp_index] == '\n')
                break;
            temp_index++;
            num_char++;
        }
        result[result_index] = malloc(sizeof(char) * (num_char + 1));
        result[result_index][0] = '\0';
        temp_result_index = 0;
        while (filecontent[index] != '\0') {
            if (filecontent[index] == '\n') {
                index++;
                break;
            }
            result[result_index][temp_result_index] = filecontent[index];
            temp_result_index++;
            index++;
        }
        result[result_index][num_char] = '\0';
        result_index++;
    }
    result[num_line + 1] = NULL;
    free(filecontent);
    return result;
}

char *load_file_in_mem(void)
{
    struct stat s;
    stat("./input.txt", &s);
    char *content = malloc(sizeof(char) * s.st_size);
    int fd = open("./input.txt", O_RDONLY);
    if (read(fd, content, s.st_size) <= 0) {
        free(content);
        close(fd);
        return NULL;
    }
    close(fd);
    return content;
}

int get_num(char *line, int index)
{
    int result = 0;
    int power = 1;
    while (line[index] != '\0') {
        if (line[index] == ' ')
            break;
        index++;
    }
    index--;
    while (index > 0) {
        if (line[index] == ' ')
            break;
        result += (line[index] - '0') * power;
        index--;
        power *= 10;
    }
    return result;
}

int main(void)
{
    int result = 1;
    char **filecontent = load_2d_arr_from_file(load_file_in_mem());
    int index_time = 5;
    int index_dist = 9;
    for (int i = 0; i < 4; i++) {
        while (filecontent[0][index_time] == ' ') {
            index_time++;
        }
        while (filecontent[1][index_dist] == ' ') {
            index_dist++;
        }
        int time = get_num(filecontent[0], index_time);
        int dist = get_num(filecontent[1], index_dist);
        printf("Time: %d  Dist: %d\n", time, dist);
        int record = 0;
        int win = 0;
        for (int i = 0; i <= time; i++) {
            record = i * (time - i);
            if (record > dist)
                win++;
        }
        result *= win;
        while (filecontent[0][index_time] != '\0') {
            if (filecontent[0][index_time] == ' ')
                break;
            index_time++;
        }
        while (filecontent[1][index_dist] != '\0') {
            if (filecontent[1][index_dist] == ' ')
                break;
            index_dist++;
        }
    }
    printf("%d\n", result);
    for (int i = 0; filecontent[i] != NULL; i++)
        free(filecontent[i]);
    free(filecontent);
    return 0;
}