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

int main(void)
{
    int result = 0;
    int temp = 0;
    int top1 = 0;
    int top2 = 0;
    int top3 = 0;
    char **filecontent = load_2d_arr_from_file(load_file_in_mem());

    for (int i = 0; filecontent[i] != NULL; i++) {
        if (strlen(filecontent[i]) != 0) {
            temp += atoi(filecontent[i]);
        } else {
            if (temp > top1) {
                top3 = top2;
                top2 = top1;
                top1 = temp;
                temp = 0;
                continue;
            }
            if (temp > top2) {
                top3 = top2;
                top2 = temp;
                temp = 0;
                continue;
            }
            if (temp > top3) {
                top3 = temp;
                temp = 0;
                continue;
            }
            temp = 0;
        }
    }
    result = top1 + top2 + top3;
    printf("Top 1 : %d  Top 2 : %d  Top 3 : %d   Total : %d\n", top1, top2, top3, result);
    for (int i = 0; filecontent[i] != NULL; i++) {
        free(filecontent[i]);
    }
    free(filecontent);
    return 0;
}