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

int is_tree_visible(char **filecontent, int i, int j)
{
    char tree = filecontent[i][j];
    int status = 0;
    for (int j1 = 0; j1 < j; j1++) {
        if (filecontent[i][j1] >= tree) {
            status += 1;
            break;
        }
    }
    for (int j1 = j + 1; filecontent[i][j1] != '\0'; j1++) {
        if (filecontent[i][j1] >= tree) {
            status += 1;
            break;
        }
    }
    for (int i1 = 0; i1 < i; i1++) {
        if (filecontent[i1][j] >= tree) {
            status += 1;
            break;
        }
    }
    for (int i1 = i + 1; filecontent[i1] != NULL; i1++) {
        if (filecontent[i1][j] >= tree) {
            status += 1;
            break;
        }
    }
    if (status == 4)
        return 0;
    return 1;
}

int main(void)
{
    char **filecontent = load_2d_arr_from_file(load_file_in_mem());
    int result = strlen(filecontent[0]);
    for (int i = 1; filecontent[i] != NULL; i++) {
        if (filecontent[i + 1] == NULL) {
            result += strlen(filecontent[i]);
            break;
        }
        result += 2;
        for (int j = 1; filecontent[i][j + 1] != '\0'; j++) {
            result += is_tree_visible(filecontent, i, j);
        }
    }
    printf("%d\n", result);
    for (int i = 0; filecontent[i] != NULL; i++)
        free(filecontent[i]);
    free(filecontent);
    return 0;
}