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

/*
A X Rock
B Y Paper
C Z Scissors
*/

int main(void)
{
    char **filecontent = load_2d_arr_from_file(load_file_in_mem());
    int my_score = 0;
    for (int i = 0; filecontent[i] != NULL; i++) {
        char op_move = filecontent[i][0];
        char my_move = filecontent[i][2];
        switch (my_move) {
            case 'X':
                my_score += 1;
                switch (op_move) {
                    case 'A':
                        my_score += 3;
                        break;
                    case 'B':
                        my_score += 0;
                        break;
                    case 'C':
                        my_score += 6;
                        break;
                }
                break;
            case 'Y':
                my_score += 2;
                switch (op_move) {
                    case 'A':
                        my_score += 6;
                        break;
                    case 'B':
                        my_score += 3;
                        break;
                    case 'C':
                        my_score += 0;
                        break;
                }
                break;
            case 'Z':
                my_score += 3;
                switch (op_move) {
                    case 'A':
                        my_score += 0;
                        break;
                    case 'B':
                        my_score += 6;
                        break;
                    case 'C':
                        my_score += 3;
                        break;
                }
                break;
        }
    }
    printf("%d\n", my_score);
    for (int i = 0; filecontent[i] != NULL; i++)
        free(filecontent[i]);
    free(filecontent);
    return 0;
}