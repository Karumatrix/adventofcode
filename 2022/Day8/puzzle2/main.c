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
    int tree_visible = 1;
    for (int j1 = j - 1; j1 >= 0; j1--) {
        if (filecontent[i][j1] >= tree) {
            tree_visible *= j - j1;
            break;
        }
        if (j1 - 1 == -1) {
            tree_visible *= j - j1;
        }
    }
    for (int j1 = j + 1; filecontent[i][j1] != '\0'; j1++) {
        if (filecontent[i][j1] >= tree) {
            tree_visible *= j1 - j;
            break;
        }
        if (filecontent[i][j1 + 1] == '\0') {
            tree_visible *= j1 - j;
        }
    }
    for (int i1 = i - 1; i1 >= 0; i1--) {
        if (filecontent[i1][j] >= tree) {
            tree_visible *= i - i1;
            break;
        }
        if (i1 - 1 == -1) {
            tree_visible *= i - i1;
        }
    }
    for (int i1 = i + 1; filecontent[i1] != NULL; i1++) {
        if (filecontent[i1][j] >= tree) {
            tree_visible *= i1 - i;
            break;
        }
        if (filecontent[i1 + 1] == NULL) {
            tree_visible *= i1 - i;
        }
    }
    return tree_visible;
}

int main(void)
{
    char **filecontent = load_2d_arr_from_file(load_file_in_mem());
    int result = 0;
    for (int i = 1; filecontent[i + 1] != NULL; i++) {
        for (int j = 1; filecontent[i][j + 1] != '\0'; j++) {
            int temp = is_tree_visible(filecontent, i, j);
            if (temp > result)
                result = temp;
        }
    }
    printf("%d\n", result);
    for (int i = 0; filecontent[i] != NULL; i++)
        free(filecontent[i]);
    free(filecontent);
    return 0;
}