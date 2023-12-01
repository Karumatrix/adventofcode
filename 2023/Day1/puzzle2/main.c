#include <fcntl.h>
#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>
#include <sys/stat.h>
#include <stddef.h>
#include <string.h>

char *nums[] = {
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
    NULL
};

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

int find_num_at(char const *line, int index)
{
    int status = 0;
    for (int j = 0; nums[j] != NULL; j++) {
        int k = 0;
        for (k; nums[j][k] != '\0'; k++) {
            if ((index + k) <= strlen(line) && line[index + k] != nums[j][k]) {
                k = 0;
                break;
            }
        }
        if (k == strlen(nums[j])) {
            return (j + 1) + '0';
        }
    }
    return -1;
}

char *find_num(char const *line)
{
    char num1 = 0;
    char t_num1 = 0;
    char num2 = 0;
    char t_num2 = 0;

    for (int i = 0; line[i] != '\0'; i++) {
        if (line[i] >= '0' && line[i] <= '9') {
            num1 = line[i];
            break;
        }
        if ((t_num1 = find_num_at(line, i)) != -1) {
            num1 = t_num1;
            break;
        }
    }
    for (int i = strlen(line); i >= 0; i--) {
        if (line[i] >= '0' && line[i] <= '9') {
            num2 = line[i];
            break;
        }
        if ((t_num2 = find_num_at(line, i)) != -1) {
            num2 = t_num2;
            break;
        }
    }
    char *result = malloc(sizeof(char) * 3);
    result[0] = num1;
    result[1] = num2;
    result[2] = '\0';
    return result;
}

int main(void)
{
    int result = 0;
    int i_num = 0;
    char **filecontent = load_2d_arr_from_file(load_file_in_mem());

    for (int i = 0; filecontent[i] != NULL; i++) {
        char *num = find_num(filecontent[i]);
        i_num = atoi(num);
        free(num);
        result += i_num;
        free(filecontent[i]);
    }
    printf("%d\n", result);
    free(filecontent);
    return 0;
}