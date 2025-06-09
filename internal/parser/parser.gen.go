package parser

import (
	_i0 "github.com/dcaiafa/lox/internal/ast"
)

var _rules = []int32{
	0, 1, 1, 2, 2, 3, 4, 4, 5, 6, 6, 7, 8, 8,
	8, 8, 9, 10, 10, 10, 10, 11, 11, 12, 13, 13, 14, 14,
	14, 14, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 24,
	24, 24, 24, 25, 25, 25, 25, 26, 26, 27, 28, 28, 29, 29,
	29, 29, 30, 31, 32, 33, 34, 34, 35, 35, 36, 36, 37, 37,
	38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44,
	45, 45, 46, 46, 47, 47, 48, 48, 49, 49, 50, 50, 51, 51,
	52, 52, 53, 53, 54, 54, 55, 55, 56, 56, 57, 57, 58, 58,
	59, 59,
}

var _termCounts = []int32{
	1, 2, 1, 1, 1, 3, 1, 1, 5, 2, 1, 2, 1, 1,
	1, 1, 6, 1, 1, 1, 1, 4, 4, 3, 1, 1, 1, 1,
	1, 1, 1, 6, 5, 4, 5, 3, 1, 1, 1, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 3, 1, 4, 1, 1, 1, 1,
	1, 1, 1, 4, 1, 4, 1, 0, 2, 1, 1, 0, 2, 1,
	1, 0, 2, 1, 1, 0, 3, 1, 2, 1, 1, 0, 1, 0,
	1, 0, 2, 1, 1, 0, 2, 1, 1, 0, 2, 1, 1, 0,
	2, 1, 2, 1, 3, 1, 2, 1, 1, 0, 1, 0, 2, 1,
	1, 0,
}

var _actions = []int32{
	157, 168, 171, 180, 187, 196, 199, 202, 205, 212, 219, 226, 229, 236,
	245, 258, 277, 284, 297, 300, 303, 316, 329, 336, 349, 352, 363, 366,
	369, 372, 393, 414, 435, 454, 473, 480, 499, 520, 539, 560, 563, 352,
	568, 583, 618, 352, 653, 656, 659, 696, 731, 742, 757, 792, 817, 842,
	845, 850, 855, 860, 873, 892, 903, 908, 911, 731, 916, 927, 930, 941,
	944, 955, 958, 969, 980, 991, 1002, 352, 1013, 1034, 352, 1039, 1064, 1089,
	1114, 1139, 1164, 1189, 1214, 1239, 1242, 1247, 1252, 1257, 1262, 1291, 1320, 1323,
	1352, 1381, 1386, 1411, 1428, 1445, 1458, 1461, 1482, 1485, 1520, 1555, 1562, 1569,
	1576, 1583, 1588, 1591, 1596, 1607, 1622, 892, 1635, 1638, 1641, 1646, 1651, 1668,
	1685, 1702, 1719, 1736, 1753, 1770, 1779, 1792, 1795, 1808, 1829, 1850, 1887, 1894,
	1897, 1900, 1903, 1906, 1909, 1912, 1931, 1942, 1953, 1958, 1770, 1971, 1974, 1977,
	1980, 1985, 1990, 10, 0, -63, 1, 1, 18, -63, 41, 2, 17, -63,
	2, 0, -2, 8, 0, -65, 18, -65, 41, -65, 17, -65, 6, 0,
	-67, 18, 6, 17, 7, 8, 0, -62, 18, -62, 41, 13, 17, -62,
	2, 0, 2147483647, 2, 41, 15, 2, 41, 14, 6, 0, -4, 18, -4,
	17, -4, 6, 0, -3, 18, -3, 17, -3, 6, 0, -69, 18, -69,
	17, -69, 2, 0, -1, 6, 0, -66, 18, 6, 17, 7, 8, 0,
	-64, 18, -64, 41, -64, 17, -64, 12, 0, -71, 34, -75, 18, -71,
	41, 17, 17, -71, 19, 18, 18, 0, -85, 32, 24, 22, 25, 34,
	26, 18, -85, 21, 27, 23, 28, 41, 29, 17, -85, 6, 0, -68,
	18, -68, 17, -68, 12, 0, -7, 34, -7, 18, -7, 41, -7, 17,
	-7, 19, -7, 2, 34, -74, 2, 34, 39, 12, 0, -6, 34, -6,
	18, -6, 41, -6, 17, -6, 19, -6, 12, 0, -73, 34, -73, 18,
	-73, 41, -73, 17, -73, 19, -73, 6, 0, -5, 18, -5, 17, -5,
	12, 0, -70, 34, -75, 18, -70, 41, 17, 17, -70, 19, 18, 2,
	34, 56, 10, 34, 43, 36, 44, 37, -109, 8, 45, 7, 46, 2,
	3, 41, 2, 34, 55, 2, 34, 40, 20, 6, -30, 0, -30, 32,
	-30, 22, -30, 34, -30, 18, -30, 21, -30, 23, -30, 41, -30, 17,
	-30, 20, 6, -29, 0, -29, 32, -29, 22, -29, 34, -29, 18, -29,
	21, -29, 23, -29, 41, -29, 17, -29, 20, 6, -27, 0, -27, 32,
	-27, 22, -27, 34, -27, 18, -27, 21, -27, 23, -27, 41, -27, 17,
	-27, 18, 0, -25, 32, -25, 22, -25, 34, -25, 18, -25, 21, -25,
	23, -25, 41, -25, 17, -25, 18, 0, -87, 32, -87, 22, -87, 34,
	-87, 18, -87, 21, -87, 23, -87, 41, -87, 17, -87, 6, 0, -23,
	18, -23, 17, -23, 18, 0, -84, 32, 24, 22, 25, 34, 26, 18,
	-84, 21, 27, 23, 28, 41, 29, 17, -84, 20, 6, -28, 0, -28,
	32, -28, 22, -28, 34, -28, 18, -28, 21, -28, 23, -28, 41, -28,
	17, -28, 18, 0, -24, 32, -24, 22, -24, 34, -24, 18, -24, 21,
	-24, 23, -24, 41, -24, 17, -24, 20, 6, -26, 0, -26, 32, -26,
	22, -26, 34, -26, 18, -26, 21, -26, 23, -26, 41, -26, 17, -26,
	2, 3, 61, 4, 41, 62, 5, -89, 14, 9, -37, 20, -37, 30,
	-37, 41, -37, 4, 80, 25, -37, 24, -37, 34, 9, -46, 20, -46,
	30, -46, 34, -46, 36, -46, 41, -46, 37, -46, 14, -46, 15, -46,
	8, -46, 4, -46, 25, -46, 24, -46, 7, -46, 12, -46, 13, -46,
	11, -46, 34, 9, -45, 20, -45, 30, -45, 34, -45, 36, -45, 41,
	-45, 37, -45, 14, -45, 15, -45, 8, -45, 4, -45, 25, -45, 24,
	-45, 7, -45, 12, -45, 13, -45, 11, -45, 2, 37, -108, 2, 37,
	91, 36, 9, -50, 20, -50, 30, -50, 34, -50, 36, -50, 41, -50,
	37, -50, 14, -50, 15, -50, 8, -50, 4, -50, 25, -50, 24, -50,
	10, 90, 7, -50, 12, -50, 13, -50, 11, -50, 34, 9, -47, 20,
	-47, 30, -47, 34, -47, 36, -47, 41, -47, 37, -47, 14, -47, 15,
	-47, 8, -47, 4, -47, 25, -47, 24, -47, 7, -47, 12, -47, 13,
	-47, 11, -47, 10, 20, 66, 30, 67, 41, -97, 25, 68, 24, 69,
	14, 9, -103, 20, -103, 30, -103, 41, -103, 4, -103, 25, -103, 24,
	-103, 34, 9, -107, 20, -107, 30, -107, 34, -107, 36, -107, 41, -107,
	37, -107, 14, 82, 15, 83, 8, -107, 4, -107, 25, -107, 24, -107,
	7, -107, 12, 84, 13, 85, 11, 86, 24, 9, -105, 20, -105, 30,
	-105, 34, -105, 36, -105, 41, -105, 37, -105, 8, -105, 4, -105, 25,
	-105, 24, -105, 7, -105, 24, 9, -38, 20, -38, 30, -38, 34, 43,
	36, 44, 41, -38, 37, -109, 8, 45, 4, -38, 25, -38, 24, -38,
	7, 46, 2, 3, 77, 4, 34, -36, 41, -36, 4, 34, -101, 41,
	-101, 4, 34, 56, 41, 78, 12, 0, -72, 34, -72, 18, -72, 41,
	-72, 17, -72, 19, -72, 18, 0, -86, 32, -86, 22, -86, 34, -86,
	18, -86, 21, -86, 23, -86, 41, -86, 17, -86, 10, 31, 93, 26,
	94, 34, 95, 28, 96, 36, 97, 4, 41, -91, 5, -91, 2, 5,
	103, 4, 41, 115, 5, -88, 10, 20, -58, 30, -58, 41, -58, 25,
	-58, 24, -58, 2, 8, 114, 10, 20, -60, 30, -60, 41, -60, 25,
	-60, 24, -60, 2, 8, 113, 10, 20, -99, 30, -99, 41, -99, 25,
	-99, 24, -99, 2, 41, 105, 10, 20, 66, 30, 67, 41, -96, 25,
	68, 24, 69, 10, 20, -54, 30, -54, 41, -54, 25, -54, 24, -54,
	10, 20, -57, 30, -57, 41, -57, 25, -57, 24, -57, 10, 20, -56,
	30, -56, 41, -56, 25, -56, 24, -56, 10, 20, -55, 30, -55, 41,
	-55, 25, -55, 24, -55, 20, 6, -35, 0, -35, 32, -35, 22, -35,
	34, -35, 18, -35, 21, -35, 23, -35, 41, -35, 17, -35, 4, 34,
	-100, 41, -100, 24, 9, -104, 20, -104, 30, -104, 34, -104, 36, -104,
	41, -104, 37, -104, 8, -104, 4, -104, 25, -104, 24, -104, 7, -104,
	24, 9, -43, 20, -43, 30, -43, 34, -43, 36, -43, 41, -43, 37,
	-43, 8, -43, 4, -43, 25, -43, 24, -43, 7, -43, 24, 9, -44,
	20, -44, 30, -44, 34, -44, 36, -44, 41, -44, 37, -44, 8, -44,
	4, -44, 25, -44, 24, -44, 7, -44, 24, 9, -41, 20, -41, 30,
	-41, 34, -41, 36, -41, 41, -41, 37, -41, 8, -41, 4, -41, 25,
	-41, 24, -41, 7, -41, 24, 9, -42, 20, -42, 30, -42, 34, -42,
	36, -42, 41, -42, 37, -42, 8, -42, 4, -42, 25, -42, 24, -42,
	7, -42, 24, 9, -40, 20, -40, 30, -40, 34, -40, 36, -40, 41,
	-40, 37, -40, 8, -40, 4, -40, 25, -40, 24, -40, 7, -40, 24,
	9, -106, 20, -106, 30, -106, 34, -106, 36, -106, 41, -106, 37, -106,
	8, -106, 4, -106, 25, -106, 24, -106, 7, -106, 24, 9, -39, 20,
	-39, 30, -39, 34, -39, 36, -39, 41, -39, 37, -39, 8, -39, 4,
	-39, 25, -39, 24, -39, 7, -39, 2, 9, 107, 4, 37, -109, 7,
	46, 4, 40, 109, 39, 110, 4, 41, 118, 4, 119, 4, 41, -10,
	4, -10, 28, 2, -14, 9, -14, 26, -14, 34, -14, 27, -14, 28,
	-14, 36, -14, 41, -14, 14, -14, 4, -14, 29, -14, 12, -14, 16,
	-14, 11, -14, 28, 2, -12, 9, -12, 26, -12, 34, -12, 27, -12,
	28, -12, 36, -12, 41, -12, 14, -12, 4, -12, 29, -12, 12, -12,
	16, -12, 11, -12, 2, 8, 131, 28, 2, -13, 9, -13, 26, -13,
	34, -13, 27, -13, 28, -13, 36, -13, 41, -13, 14, -13, 4, -13,
	29, -13, 12, -13, 16, -13, 11, -13, 28, 2, -15, 9, -15, 26,
	-15, 34, -15, 27, -15, 28, -15, 36, -15, 41, -15, 14, -15, 4,
	-15, 29, -15, 12, -15, 16, -15, 11, -15, 4, 41, -77, 4, -77,
	24, 26, -83, 34, -83, 27, -83, 28, -83, 36, -83, 41, -83, 14,
	125, 4, -83, 29, -83, 12, 126, 16, 127, 11, 128, 16, 26, -79,
	34, -79, 27, -79, 28, -79, 36, -79, 41, -79, 4, -79, 29, -79,
	16, 26, 94, 34, 95, 27, 120, 28, 96, 36, 97, 41, -81, 4,
	-81, 29, 121, 12, 6, -93, 32, 24, 22, 25, 34, 26, 21, 27,
	41, 29, 2, 41, 135, 20, 6, -33, 0, -33, 32, -33, 22, -33,
	34, -33, 18, -33, 21, -33, 23, -33, 41, -33, 17, -33, 2, 41,
	136, 34, 9, -48, 20, -48, 30, -48, 34, -48, 36, -48, 41, -48,
	37, -48, 14, -48, 15, -48, 8, -48, 4, -48, 25, -48, 24, -48,
	7, -48, 12, -48, 13, -48, 11, -48, 34, 9, -49, 20, -49, 30,
	-49, 34, -49, 36, -49, 41, -49, 37, -49, 14, -49, 15, -49, 8,
	-49, 4, -49, 25, -49, 24, -49, 7, -49, 12, -49, 13, -49, 11,
	-49, 6, 38, -52, 40, -52, 39, -52, 6, 38, -53, 40, -53, 39,
	-53, 6, 38, -111, 40, -111, 39, -111, 6, 38, 137, 40, 109, 39,
	110, 4, 9, -113, 34, 139, 2, 34, 141, 4, 41, -90, 5, -90,
	10, 20, -98, 30, -98, 41, -98, 25, -98, 24, -98, 14, 9, -102,
	20, -102, 30, -102, 41, -102, 4, -102, 25, -102, 24, -102, 12, 0,
	-8, 34, -8, 18, -8, 41, -8, 17, -8, 19, -8, 2, 8, 143,
	2, 8, 144, 4, 41, -80, 4, -80, 4, 41, -9, 4, -9, 16,
	26, -78, 34, -78, 27, -78, 28, -78, 36, -78, 41, -78, 4, -78,
	29, -78, 16, 26, -19, 34, -19, 27, -19, 28, -19, 36, -19, 41,
	-19, 4, -19, 29, -19, 16, 26, -17, 34, -17, 27, -17, 28, -17,
	36, -17, 41, -17, 4, -17, 29, -17, 16, 26, -18, 34, -18, 27,
	-18, 28, -18, 36, -18, 41, -18, 4, -18, 29, -18, 16, 26, -20,
	34, -20, 27, -20, 28, -20, 36, -20, 41, -20, 4, -20, 29, -20,
	16, 26, -82, 34, -82, 27, -82, 28, -82, 36, -82, 41, -82, 4,
	-82, 29, -82, 16, 26, -11, 34, -11, 27, -11, 28, -11, 36, -11,
	41, -11, 4, -11, 29, -11, 8, 26, 94, 34, 95, 28, 96, 36,
	97, 12, 6, -95, 32, -95, 22, -95, 34, -95, 21, -95, 41, -95,
	2, 6, 145, 12, 6, -92, 32, 24, 22, 25, 34, 26, 21, 27,
	41, 29, 20, 6, -32, 0, -32, 32, -32, 22, -32, 34, -32, 18,
	-32, 21, -32, 23, -32, 41, -32, 17, -32, 20, 6, -34, 0, -34,
	32, -34, 22, -34, 34, -34, 18, -34, 21, -34, 23, -34, 41, -34,
	17, -34, 36, 9, -51, 20, -51, 30, -51, 34, -51, 36, -51, 41,
	-51, 37, -51, 14, -51, 15, -51, 8, -51, 4, -51, 25, -51, 24,
	-51, 10, -51, 7, -51, 12, -51, 13, -51, 11, -51, 6, 38, -110,
	40, -110, 39, -110, 2, 9, -112, 2, 9, 146, 2, 9, 147, 2,
	2, 150, 2, 35, 151, 2, 35, 152, 18, 0, -31, 32, -31, 22,
	-31, 34, -31, 18, -31, 21, -31, 23, -31, 41, -31, 17, -31, 10,
	20, -59, 30, -59, 41, -59, 25, -59, 24, -59, 10, 20, -61, 30,
	-61, 41, -61, 25, -61, 24, -61, 4, 41, -76, 4, -76, 12, 6,
	-94, 32, -94, 22, -94, 34, -94, 21, -94, 41, -94, 2, 9, 154,
	2, 9, 155, 2, 9, 156, 4, 41, -21, 4, -21, 4, 41, -22,
	4, -22, 28, 2, -16, 9, -16, 26, -16, 34, -16, 27, -16, 28,
	-16, 36, -16, 41, -16, 14, -16, 4, -16, 29, -16, 12, -16, 16,
	-16, 11, -16,
}

var _goto = []int32{
	157, 164, 164, 165, 164, 164, 164, 164, 164, 164, 164, 164, 176, 164,
	183, 194, 164, 164, 164, 164, 164, 164, 164, 213, 220, 225, 164, 164,
	164, 164, 164, 164, 164, 164, 164, 244, 164, 164, 164, 164, 259, 264,
	164, 164, 164, 283, 164, 164, 164, 164, 302, 164, 317, 164, 322, 164,
	164, 164, 333, 164, 164, 336, 164, 164, 164, 349, 164, 164, 164, 164,
	164, 164, 364, 164, 164, 164, 164, 375, 164, 164, 394, 164, 164, 164,
	164, 164, 164, 164, 164, 164, 409, 414, 164, 164, 164, 164, 164, 164,
	164, 164, 419, 164, 424, 435, 164, 164, 164, 164, 164, 164, 164, 164,
	450, 453, 164, 164, 164, 164, 164, 456, 164, 164, 164, 164, 164, 164,
	164, 164, 164, 164, 164, 467, 164, 164, 472, 164, 164, 164, 164, 164,
	164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 483, 164, 164, 164,
	164, 164, 164, 6, 34, 3, 35, 4, 1, 5, 0, 10, 12, 8,
	3, 9, 2, 10, 36, 11, 37, 12, 6, 12, 8, 3, 9, 2,
	16, 10, 40, 19, 5, 20, 4, 21, 38, 22, 39, 23, 18, 19,
	30, 17, 31, 14, 32, 13, 33, 45, 34, 46, 35, 18, 36, 15,
	37, 16, 38, 6, 40, 19, 5, 20, 4, 59, 4, 20, 57, 53,
	58, 18, 54, 42, 57, 47, 27, 48, 26, 49, 21, 50, 22, 51,
	25, 52, 23, 53, 55, 54, 14, 19, 30, 17, 31, 14, 32, 13,
	60, 18, 36, 15, 37, 16, 38, 4, 47, 63, 48, 64, 18, 54,
	42, 57, 47, 27, 48, 26, 49, 21, 65, 22, 51, 25, 52, 23,
	53, 55, 54, 18, 54, 42, 57, 47, 27, 48, 26, 49, 21, 89,
	22, 51, 25, 52, 23, 53, 55, 54, 14, 29, 70, 51, 71, 52,
	72, 30, 73, 33, 74, 32, 75, 31, 76, 4, 24, 87, 56, 88,
	10, 57, 47, 27, 48, 26, 49, 25, 52, 23, 81, 2, 20, 79,
	12, 41, 92, 9, 98, 6, 99, 8, 100, 7, 101, 42, 102, 14,
	29, 70, 51, 104, 52, 72, 30, 73, 33, 74, 32, 75, 31, 76,
	10, 29, 116, 30, 73, 33, 74, 32, 75, 31, 76, 18, 54, 42,
	57, 47, 27, 48, 26, 49, 21, 106, 22, 51, 25, 52, 23, 53,
	55, 54, 14, 57, 47, 27, 48, 26, 49, 22, 117, 25, 52, 23,
	53, 55, 54, 4, 57, 47, 27, 108, 4, 28, 111, 58, 112, 4,
	10, 129, 44, 130, 10, 9, 98, 11, 122, 43, 123, 8, 100, 7,
	124, 14, 19, 30, 17, 31, 14, 132, 49, 133, 50, 134, 18, 36,
	16, 38, 2, 28, 138, 2, 59, 140, 10, 9, 98, 6, 148, 8,
	100, 7, 101, 42, 102, 4, 9, 98, 8, 142, 10, 19, 30, 17,
	31, 14, 149, 18, 36, 16, 38, 4, 9, 98, 8, 153,
}

type _Bounds struct {
	Begin Token
	End   Token
	Empty bool
}

func _cast[T any](v any) T {
	cv, _ := v.(T)
	return cv
}

type Error struct {
	Token    Token
	Expected []int
}

func _Find(table []int32, y, x int32) (int32, bool) {
	i := int(table[int(y)])
	count := int(table[i])
	i++
	end := i + count
	for ; i < end; i += 2 {
		if table[i] == x {
			return table[i+1], true
		}
	}
	return 0, false
}

type _Lexer interface {
	ReadToken() (Token, int)
}

type _item struct {
	State  int32
	Sym    any
	Bounds _Bounds
}

type lox struct {
	_lex   _Lexer
	_stack _Stack[_item]

	_la    int
	_lasym any

	_qla    int
	_qlasym any
}

func (p *parser) parse(lex _Lexer) bool {
	const accept = 2147483647

	p._lex = lex
	p._qla = -1
	p._stack.Push(_item{})

	p._readToken()

	for {
		topState := p._stack.Peek(0).State
		action, ok := _Find(_actions, topState, int32(p._la))
		if !ok {
			if !p._recover() {
				return false
			}
			continue
		}
		if action == accept {
			break
		} else if action >= 0 { // shift
			latok, ok := p._lasym.(Token)
			if !ok {
				latok = p._lasym.(Error).Token
			}
			p._stack.Push(_item{
				State: action,
				Sym:   p._lasym,
				Bounds: _Bounds{
					Begin: latok,
					End:   latok,
				},
			})
			p._readToken()
		} else { // reduce
			prod := -action
			termCount := _termCounts[int(prod)]
			rule := _rules[int(prod)]
			res := p._act(prod)

			// Compute reduction token bounds.
			// Trim leading and trailing empty bounds.
			boundSlice := p._stack.PeekSlice(int(termCount))
			for len(boundSlice) > 0 && boundSlice[0].Bounds.Empty {
				boundSlice = boundSlice[1:]
			}
			for len(boundSlice) > 0 && boundSlice[len(boundSlice)-1].Bounds.Empty {
				boundSlice = boundSlice[:len(boundSlice)-1]
			}
			var bounds _Bounds
			if len(boundSlice) > 0 {
				bounds.Begin = boundSlice[0].Bounds.Begin
				bounds.End = boundSlice[len(boundSlice)-1].Bounds.End
			} else {
				bounds.Empty = true
			}
			if !bounds.Empty {
				p._onBounds(res, bounds.Begin, bounds.End)
			}
			p._stack.Pop(int(termCount))
			topState = p._stack.Peek(0).State
			nextState, _ := _Find(_goto, topState, rule)
			p._stack.Push(_item{
				State:  nextState,
				Sym:    res,
				Bounds: bounds,
			})
		}
	}

	return true
}

// recoverLookahead can be called during an error production action (an action
// for a production that has a @error term) to recover the lookahead that was
// possibly lost in the process of reducing the error production.
func (p *parser) recoverLookahead(typ int, tok Token) {
	if p._qla != -1 {
		panic("recovered lookahead already pending")
	}

	p._qla = p._la
	p._qlasym = p._lasym
	p._la = typ
	p._lasym = tok
}

func (p *parser) _readToken() {
	if p._qla != -1 {
		p._la = p._qla
		p._lasym = p._qlasym
		p._qla = -1
		p._qlasym = nil
		return
	}

	p._lasym, p._la = p._lex.ReadToken()
	if p._la == ERROR {
		p._lasym = p._makeError()
	}
}

func (p *parser) _recover() bool {
	errSym, ok := p._lasym.(Error)
	if !ok {
		errSym = p._makeError()
	}

	for p._la == ERROR {
		p._readToken()
	}

	for {
		save := p._stack

		for len(p._stack) >= 1 {
			state := p._stack.Peek(0).State

			for {
				action, ok := _Find(_actions, state, int32(ERROR))
				if !ok {
					break
				}

				if action < 0 {
					prod := -action
					rule := _rules[int(prod)]
					state, _ = _Find(_goto, state, rule)
					continue
				}

				state = action

				_, ok = _Find(_actions, state, int32(p._la))
				if !ok {
					break
				}

				p._qla = p._la
				p._qlasym = p._lasym
				p._la = ERROR
				p._lasym = errSym
				return true
			}

			p._stack.Pop(1)
		}

		if p._la == EOF {
			return false
		}

		p._stack = save
		p._readToken()
	}
}

func (p *parser) _makeError() Error {
	e := Error{
		Token: p._lasym.(Token),
	}

	// Compile list of allowed tokens at this state.
	// See _Find for the format of the _actions table.
	s := p._stack.Peek(0).State
	i := int(_actions[int(s)])
	count := int(_actions[i])
	i++
	end := i + count
	for ; i < end; i += 2 {
		e.Expected = append(e.Expected, int(_actions[i]))
	}

	return e
}

func (p *parser) _act(prod int32) any {
	switch prod {
	case 1:
		return p.on_spec(
			_cast[[]Token](p._stack.Peek(1).Sym),
			_cast[[][]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 2:
		return p.on_spec__error(
			_cast[Error](p._stack.Peek(0).Sym),
		)
	case 3:
		return p.on_section(
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 4:
		return p.on_section(
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 5:
		return p.on_parser_section(
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 6:
		return p.on_parser_statement(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 7:
		return p.on_parser_statement__nl(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 8:
		return p.on_parser_rule(
			_cast[Token](p._stack.Peek(4).Sym),
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[[]*_i0.ParserProd](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 9:
		return p.on_parser_prod(
			_cast[[]*_i0.ParserTerm](p._stack.Peek(1).Sym),
			_cast[*_i0.ProdQualifier](p._stack.Peek(0).Sym),
		)
	case 10:
		return p.on_parser_prod__empty(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 11:
		return p.on_parser_term_card(
			_cast[*_i0.ParserTerm](p._stack.Peek(1).Sym),
			_cast[_i0.ParserTermType](p._stack.Peek(0).Sym),
		)
	case 12:
		return p.on_parser_term__token(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 13:
		return p.on_parser_term__token(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 14:
		return p.on_parser_term__token(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 15:
		return p.on_parser_term__list(
			_cast[*_i0.ParserTerm](p._stack.Peek(0).Sym),
		)
	case 16:
		return p.on_parser_list(
			_cast[Token](p._stack.Peek(5).Sym),
			_cast[Token](p._stack.Peek(4).Sym),
			_cast[*_i0.ParserTerm](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[*_i0.ParserTerm](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 17:
		return p.on_parser_card(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 18:
		return p.on_parser_card(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 19:
		return p.on_parser_card(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 20:
		return p.on_parser_card(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 21:
		return p.on_parser_qualif(
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 22:
		return p.on_parser_qualif(
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 23:
		return p.on_lexer_section(
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 24:
		return p.on_lexer_statement(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 25:
		return p.on_lexer_statement(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 26:
		return p.on_lexer_rule(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 27:
		return p.on_lexer_rule(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 28:
		return p.on_lexer_rule(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 29:
		return p.on_lexer_rule(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 30:
		return p.on_lexer_rule__nl(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 31:
		return p.on_mode(
			_cast[Token](p._stack.Peek(5).Sym),
			_cast[Token](p._stack.Peek(4).Sym),
			_cast[[]Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[[]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 32:
		return p.on_token_rule(
			_cast[Token](p._stack.Peek(4).Sym),
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[*_i0.LexerExpr](p._stack.Peek(2).Sym),
			_cast[[]_i0.Action](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 33:
		return p.on_frag_rule(
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[*_i0.LexerExpr](p._stack.Peek(2).Sym),
			_cast[[]_i0.Action](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 34:
		return p.on_macro_rule(
			_cast[Token](p._stack.Peek(4).Sym),
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[*_i0.LexerExpr](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 35:
		return p.on_external_rule(
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[[]*_i0.ExternalName](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 36:
		return p.on_external_name(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 37:
		return p.on_lexer_expr(
			_cast[[]*_i0.LexerFactor](p._stack.Peek(0).Sym),
		)
	case 38:
		return p.on_lexer_factor(
			_cast[[]*_i0.LexerTermCard](p._stack.Peek(0).Sym),
		)
	case 39:
		return p.on_lexer_term_card(
			_cast[_i0.LexerTerm](p._stack.Peek(1).Sym),
			_cast[_i0.Card](p._stack.Peek(0).Sym),
		)
	case 40:
		return p.on_lexer_card(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 41:
		return p.on_lexer_card(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 42:
		return p.on_lexer_card(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 43:
		return p.on_lexer_card(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 44:
		return p.on_lexer_card(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 45:
		return p.on_lexer_term__tok(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 46:
		return p.on_lexer_term__tok(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 47:
		return p.on_lexer_term__char_class_expr(
			_cast[_i0.CharClassExpr](p._stack.Peek(0).Sym),
		)
	case 48:
		return p.on_lexer_term__expr(
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[*_i0.LexerExpr](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 49:
		return p.on_char_class_expr__binary(
			_cast[_i0.CharClassExpr](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[_i0.CharClassExpr](p._stack.Peek(0).Sym),
		)
	case 50:
		return p.on_char_class_expr__char_class(
			_cast[*_i0.CharClass](p._stack.Peek(0).Sym),
		)
	case 51:
		return p.on_char_class(
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[[]Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 52:
		return p.on_char_class_item(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 53:
		return p.on_char_class_item(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 54:
		return p.on_action(
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 55:
		return p.on_action(
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 56:
		return p.on_action(
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 57:
		return p.on_action(
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 58:
		return p.on_action_discard(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 59:
		return p.on_action_push_mode(
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 60:
		return p.on_action_pop_mode(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 61:
		return p.on_action_emit(
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 62: // ZeroOrMore
		return _cast[[]Token](p._stack.Peek(0).Sym)
	case 63: // ZeroOrMore
		{
			var zero []Token
			return zero
		}
	case 64:
		{ // OneOrMoreF
			l := _cast[[]Token](p._stack.Peek(1).Sym)
			e := _cast[Token](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 65:
		{ // OneOrMoreF
			var l []Token
			e := _cast[Token](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 66: // ZeroOrMore
		return _cast[[][]_i0.Statement](p._stack.Peek(0).Sym)
	case 67: // ZeroOrMore
		{
			var zero [][]_i0.Statement
			return zero
		}
	case 68: // OneOrMore
		return append(
			_cast[[][]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 69: // OneOrMore
		return [][]_i0.Statement{
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		}
	case 70: // ZeroOrMore
		return _cast[[]_i0.Statement](p._stack.Peek(0).Sym)
	case 71: // ZeroOrMore
		{
			var zero []_i0.Statement
			return zero
		}
	case 72: // OneOrMore
		return append(
			_cast[[]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 73: // OneOrMore
		return []_i0.Statement{
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		}
	case 74: // ZeroOrOne
		return _cast[Token](p._stack.Peek(0).Sym)
	case 75: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	case 76: // List
		return append(
			_cast[[]*_i0.ParserProd](p._stack.Peek(2).Sym),
			_cast[*_i0.ParserProd](p._stack.Peek(0).Sym),
		)
	case 77: // List
		return []*_i0.ParserProd{
			_cast[*_i0.ParserProd](p._stack.Peek(0).Sym),
		}
	case 78: // OneOrMore
		return append(
			_cast[[]*_i0.ParserTerm](p._stack.Peek(1).Sym),
			_cast[*_i0.ParserTerm](p._stack.Peek(0).Sym),
		)
	case 79: // OneOrMore
		return []*_i0.ParserTerm{
			_cast[*_i0.ParserTerm](p._stack.Peek(0).Sym),
		}
	case 80: // ZeroOrOne
		return _cast[*_i0.ProdQualifier](p._stack.Peek(0).Sym)
	case 81: // ZeroOrOne
		{
			var zero *_i0.ProdQualifier
			return zero
		}
	case 82: // ZeroOrOne
		return _cast[_i0.ParserTermType](p._stack.Peek(0).Sym)
	case 83: // ZeroOrOne
		{
			var zero _i0.ParserTermType
			return zero
		}
	case 84: // ZeroOrMore
		return _cast[[]_i0.Statement](p._stack.Peek(0).Sym)
	case 85: // ZeroOrMore
		{
			var zero []_i0.Statement
			return zero
		}
	case 86:
		{ // OneOrMoreF
			l := _cast[[]_i0.Statement](p._stack.Peek(1).Sym)
			e := _cast[_i0.Statement](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 87:
		{ // OneOrMoreF
			var l []_i0.Statement
			e := _cast[_i0.Statement](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 88: // ZeroOrMore
		return _cast[[]Token](p._stack.Peek(0).Sym)
	case 89: // ZeroOrMore
		{
			var zero []Token
			return zero
		}
	case 90: // OneOrMore
		return append(
			_cast[[]Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 91: // OneOrMore
		return []Token{
			_cast[Token](p._stack.Peek(0).Sym),
		}
	case 92: // ZeroOrMore
		return _cast[[]_i0.Statement](p._stack.Peek(0).Sym)
	case 93: // ZeroOrMore
		{
			var zero []_i0.Statement
			return zero
		}
	case 94: // OneOrMore
		return append(
			_cast[[]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 95: // OneOrMore
		return []_i0.Statement{
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		}
	case 96: // ZeroOrMore
		return _cast[[]_i0.Action](p._stack.Peek(0).Sym)
	case 97: // ZeroOrMore
		{
			var zero []_i0.Action
			return zero
		}
	case 98: // OneOrMore
		return append(
			_cast[[]_i0.Action](p._stack.Peek(1).Sym),
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 99: // OneOrMore
		return []_i0.Action{
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		}
	case 100: // OneOrMore
		return append(
			_cast[[]*_i0.ExternalName](p._stack.Peek(1).Sym),
			_cast[*_i0.ExternalName](p._stack.Peek(0).Sym),
		)
	case 101: // OneOrMore
		return []*_i0.ExternalName{
			_cast[*_i0.ExternalName](p._stack.Peek(0).Sym),
		}
	case 102: // List
		return append(
			_cast[[]*_i0.LexerFactor](p._stack.Peek(2).Sym),
			_cast[*_i0.LexerFactor](p._stack.Peek(0).Sym),
		)
	case 103: // List
		return []*_i0.LexerFactor{
			_cast[*_i0.LexerFactor](p._stack.Peek(0).Sym),
		}
	case 104: // OneOrMore
		return append(
			_cast[[]*_i0.LexerTermCard](p._stack.Peek(1).Sym),
			_cast[*_i0.LexerTermCard](p._stack.Peek(0).Sym),
		)
	case 105: // OneOrMore
		return []*_i0.LexerTermCard{
			_cast[*_i0.LexerTermCard](p._stack.Peek(0).Sym),
		}
	case 106: // ZeroOrOne
		return _cast[_i0.Card](p._stack.Peek(0).Sym)
	case 107: // ZeroOrOne
		{
			var zero _i0.Card
			return zero
		}
	case 108: // ZeroOrOne
		return _cast[Token](p._stack.Peek(0).Sym)
	case 109: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	case 110: // OneOrMore
		return append(
			_cast[[]Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 111: // OneOrMore
		return []Token{
			_cast[Token](p._stack.Peek(0).Sym),
		}
	case 112: // ZeroOrOne
		return _cast[Token](p._stack.Peek(0).Sym)
	case 113: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	default:
		panic("unreachable")
	}
}
