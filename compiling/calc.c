#include <stdio.h>
#include <stdlib.h>

double expression(void);

double vars[26]; // variables

// 读取一个字节
char get(void) { // get one byte
    char c = getchar();
    return c;
}

// 预读一个字节，读出后再把字节退回到输入中
char peek(void) { // peek at next byte
    char c = getchar();
    ungetc(c, stdin);
    return c;
}

// 读取一个双精度浮点数
double number(void) { // read one double
    double d;
    scanf("%lf", &d);
    return d;
}

// 判断是否为期望值
void expect(char c) { // expect char c from stream
    char d = get();
    if (c != d) {
        fprintf(stderr, "Error: Expected %c but got %c.\n", c, d);
    }
}

double factor(void) { // read a factor
    double f;
    char c = peek();
    if (c == '(') { // an expression inside parantesis?
        expect('(');
        f = expression();
        expect(')');
    } else if (c >= 'A' && c <= 'Z') { // a variable ?
        expect(c);
        f = vars[c - 'A'];
    } else { // or, a number?
        f = number();
    }
    return f;
}

double term(void) { // read a term
    double t = factor();
    while (peek() == '*' || peek() == '/') { // * or / more factors
        char c = get();
        if (c == '*') {
            t = t * factor();
        } else {
            t = t / factor();
        }
    }
    return t;
}

double expression(void) { // read an expression
    double e = term();
    while (peek() == '+' || peek() == '-') { // + or - more terms
        char c = get();
        if (c == '+') {
            e = e + term();
        } else {
            e = e - term();
        }
    }
    return e;
}

double statement(void) { // read a statement
    double ret;
    char c = peek();
    // 变量只能使用 A-Z 单个字节
    if (c >= 'A' && c <= 'Z') { // variable ?
        expect(c); // 预读一个字节
        if (peek() == '=') { // assignment ?
            expect('=');
            double val = expression();
            vars[c - 'A'] = val;
            ret = val;
        } else {
            ungetc(c, stdin);
            ret = expression();
        }
    } else {
        ret = expression();
    }
    expect('\n');
    return ret;
}

int main(void) {
    printf("> "); fflush(stdout);

    for (;;) {
        double v = statement();
        printf(" = %lf\n> ", v); fflush(stdout);
    }=
    return EXIT_SUCCESS;
}