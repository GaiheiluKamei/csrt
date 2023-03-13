# csrt

## Explain

An example of an incorrectly formatted SRT file in the classical style is as follows:

```srt
0
1
00:00:00,540 --> 00:00:07,440
Hello and welcome to this video where we will be explaining the problem from Codility called equilibrium.
1

2
00:00:08,590 --> 00:00:09,440
In this problem
2

3
00:00:09,460 --> 00:00:18,430
we are given an input array containing n items and the aim of this problem is to find the point in
3

...
```

`csrt` is a tool to correct this kind of files.

> Timestamp format for SRT files: `hours:minutes:seconds,milliseconds`.

## Usage

```bash
go run . -f "incorrect.srt" -d "rn" -o "correct.srt"
```