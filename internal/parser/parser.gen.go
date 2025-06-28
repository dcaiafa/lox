package parser

import (
	_i0 "github.com/dcaiafa/lox/internal/ast"
)

var _rules = []int32{
	0, 1, 1, 2, 2, 3, 4, 4, 5, 6, 6, 7, 8, 8,
	8, 8, 9, 10, 10, 10, 10, 11, 11, 12, 13, 13, 14, 14,
	14, 14, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 24,
	24, 24, 24, 25, 25, 25, 25, 25, 26, 26, 27, 28, 28, 29,
	29, 29, 29, 30, 31, 32, 33, 34, 34, 35, 35, 36, 36, 37,
	37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44,
	44, 45, 45, 46, 46, 47, 47, 48, 48, 49, 49, 50, 50, 51,
	51, 52, 52, 53, 53, 54, 54, 55, 55, 56, 56, 57, 57, 58,
	58, 59, 59,
}

var _termCounts = []int32{
	1, 2, 1, 1, 1, 3, 1, 1, 5, 2, 1, 2, 1, 1,
	1, 1, 6, 1, 1, 1, 1, 4, 4, 3, 1, 1, 1, 1,
	1, 1, 1, 6, 5, 4, 5, 3, 1, 1, 1, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 3, 1, 4, 1, 1, 1,
	1, 1, 1, 1, 4, 1, 4, 1, 0, 2, 1, 1, 0, 2,
	1, 1, 0, 2, 1, 1, 0, 3, 1, 2, 1, 1, 0, 1,
	0, 1, 0, 2, 1, 1, 0, 2, 1, 1, 0, 2, 1, 1,
	0, 2, 1, 2, 1, 3, 1, 2, 1, 1, 0, 1, 0, 2,
	1, 1, 0,
}

var _actions = []int32{
	158, 169, 172, 181, 188, 197, 200, 203, 206, 213, 220, 227, 230, 237,
	246, 259, 278, 285, 298, 301, 304, 317, 330, 337, 350, 353, 366, 369,
	372, 375, 396, 417, 438, 457, 476, 483, 502, 523, 542, 563, 566, 353,
	571, 586, 623, 660, 353, 697, 700, 703, 742, 779, 790, 805, 842, 869,
	896, 899, 904, 909, 914, 927, 946, 957, 962, 965, 779, 970, 981, 984,
	995, 998, 1009, 1012, 1023, 1034, 1045, 1056, 353, 1067, 1088, 353, 1093, 1120,
	1147, 1174, 1201, 1228, 1255, 1282, 1309, 1312, 1317, 1322, 1327, 1332, 1361, 1390,
	1393, 1422, 1451, 1456, 1481, 1498, 1515, 1528, 1531, 1552, 1555, 1592, 1629, 1636,
	1643, 1650, 1657, 1662, 1665, 1670, 1681, 1696, 946, 1709, 1712, 1715, 1720, 1725,
	1742, 1759, 1776, 1793, 1810, 1827, 1844, 1853, 1866, 1869, 1882, 1903, 1924, 1963,
	1970, 1973, 1976, 1979, 1982, 1985, 1988, 2007, 2018, 2029, 2034, 1844, 2047, 2050,
	2053, 2056, 2061, 2066, 10, 0, -64, 1, 1, 19, -64, 42, 2, 18,
	-64, 2, 0, -2, 8, 0, -66, 19, -66, 42, -66, 18, -66, 6,
	0, -68, 19, 6, 18, 7, 8, 0, -63, 19, -63, 42, 13, 18,
	-63, 2, 0, 2147483647, 2, 42, 15, 2, 42, 14, 6, 0, -4, 19,
	-4, 18, -4, 6, 0, -3, 19, -3, 18, -3, 6, 0, -70, 19,
	-70, 18, -70, 2, 0, -1, 6, 0, -67, 19, 6, 18, 7, 8,
	0, -65, 19, -65, 42, -65, 18, -65, 12, 0, -72, 35, -76, 19,
	-72, 42, 17, 18, -72, 20, 18, 18, 0, -86, 33, 24, 23, 25,
	35, 26, 19, -86, 22, 27, 24, 28, 42, 29, 18, -86, 6, 0,
	-69, 19, -69, 18, -69, 12, 0, -7, 35, -7, 19, -7, 42, -7,
	18, -7, 20, -7, 2, 35, -75, 2, 35, 39, 12, 0, -6, 35,
	-6, 19, -6, 42, -6, 18, -6, 20, -6, 12, 0, -74, 35, -74,
	19, -74, 42, -74, 18, -74, 20, -74, 6, 0, -5, 19, -5, 18,
	-5, 12, 0, -71, 35, -76, 19, -71, 42, 17, 18, -71, 20, 18,
	2, 35, 57, 12, 11, 43, 35, 44, 37, 45, 38, -110, 8, 46,
	7, 47, 2, 3, 41, 2, 35, 56, 2, 35, 40, 20, 6, -30,
	0, -30, 33, -30, 23, -30, 35, -30, 19, -30, 22, -30, 24, -30,
	42, -30, 18, -30, 20, 6, -29, 0, -29, 33, -29, 23, -29, 35,
	-29, 19, -29, 22, -29, 24, -29, 42, -29, 18, -29, 20, 6, -27,
	0, -27, 33, -27, 23, -27, 35, -27, 19, -27, 22, -27, 24, -27,
	42, -27, 18, -27, 18, 0, -25, 33, -25, 23, -25, 35, -25, 19,
	-25, 22, -25, 24, -25, 42, -25, 18, -25, 18, 0, -88, 33, -88,
	23, -88, 35, -88, 19, -88, 22, -88, 24, -88, 42, -88, 18, -88,
	6, 0, -23, 19, -23, 18, -23, 18, 0, -85, 33, 24, 23, 25,
	35, 26, 19, -85, 22, 27, 24, 28, 42, 29, 18, -85, 20, 6,
	-28, 0, -28, 33, -28, 23, -28, 35, -28, 19, -28, 22, -28, 24,
	-28, 42, -28, 18, -28, 18, 0, -24, 33, -24, 23, -24, 35, -24,
	19, -24, 22, -24, 24, -24, 42, -24, 18, -24, 20, 6, -26, 0,
	-26, 33, -26, 23, -26, 35, -26, 19, -26, 22, -26, 24, -26, 42,
	-26, 18, -26, 2, 3, 62, 4, 42, 63, 5, -90, 14, 9, -37,
	21, -37, 31, -37, 42, -37, 4, 81, 26, -37, 25, -37, 36, 9,
	-47, 21, -47, 11, -47, 31, -47, 35, -47, 37, -47, 42, -47, 38,
	-47, 15, -47, 16, -47, 8, -47, 4, -47, 26, -47, 25, -47, 7,
	-47, 13, -47, 14, -47, 12, -47, 36, 9, -46, 21, -46, 11, -46,
	31, -46, 35, -46, 37, -46, 42, -46, 38, -46, 15, -46, 16, -46,
	8, -46, 4, -46, 26, -46, 25, -46, 7, -46, 13, -46, 14, -46,
	12, -46, 36, 9, -45, 21, -45, 11, -45, 31, -45, 35, -45, 37,
	-45, 42, -45, 38, -45, 15, -45, 16, -45, 8, -45, 4, -45, 26,
	-45, 25, -45, 7, -45, 13, -45, 14, -45, 12, -45, 2, 38, -109,
	2, 38, 92, 38, 9, -51, 21, -51, 11, -51, 31, -51, 35, -51,
	37, -51, 42, -51, 38, -51, 15, -51, 16, -51, 8, -51, 4, -51,
	26, -51, 25, -51, 10, 91, 7, -51, 13, -51, 14, -51, 12, -51,
	36, 9, -48, 21, -48, 11, -48, 31, -48, 35, -48, 37, -48, 42,
	-48, 38, -48, 15, -48, 16, -48, 8, -48, 4, -48, 26, -48, 25,
	-48, 7, -48, 13, -48, 14, -48, 12, -48, 10, 21, 67, 31, 68,
	42, -98, 26, 69, 25, 70, 14, 9, -104, 21, -104, 31, -104, 42,
	-104, 4, -104, 26, -104, 25, -104, 36, 9, -108, 21, -108, 11, -108,
	31, -108, 35, -108, 37, -108, 42, -108, 38, -108, 15, 83, 16, 84,
	8, -108, 4, -108, 26, -108, 25, -108, 7, -108, 13, 85, 14, 86,
	12, 87, 26, 9, -106, 21, -106, 11, -106, 31, -106, 35, -106, 37,
	-106, 42, -106, 38, -106, 8, -106, 4, -106, 26, -106, 25, -106, 7,
	-106, 26, 9, -38, 21, -38, 11, 43, 31, -38, 35, 44, 37, 45,
	42, -38, 38, -110, 8, 46, 4, -38, 26, -38, 25, -38, 7, 47,
	2, 3, 78, 4, 35, -36, 42, -36, 4, 35, -102, 42, -102, 4,
	35, 57, 42, 79, 12, 0, -73, 35, -73, 19, -73, 42, -73, 18,
	-73, 20, -73, 18, 0, -87, 33, -87, 23, -87, 35, -87, 19, -87,
	22, -87, 24, -87, 42, -87, 18, -87, 10, 32, 94, 27, 95, 35,
	96, 29, 97, 37, 98, 4, 42, -92, 5, -92, 2, 5, 104, 4,
	42, 116, 5, -89, 10, 21, -59, 31, -59, 42, -59, 26, -59, 25,
	-59, 2, 8, 115, 10, 21, -61, 31, -61, 42, -61, 26, -61, 25,
	-61, 2, 8, 114, 10, 21, -100, 31, -100, 42, -100, 26, -100, 25,
	-100, 2, 42, 106, 10, 21, 67, 31, 68, 42, -97, 26, 69, 25,
	70, 10, 21, -55, 31, -55, 42, -55, 26, -55, 25, -55, 10, 21,
	-58, 31, -58, 42, -58, 26, -58, 25, -58, 10, 21, -57, 31, -57,
	42, -57, 26, -57, 25, -57, 10, 21, -56, 31, -56, 42, -56, 26,
	-56, 25, -56, 20, 6, -35, 0, -35, 33, -35, 23, -35, 35, -35,
	19, -35, 22, -35, 24, -35, 42, -35, 18, -35, 4, 35, -101, 42,
	-101, 26, 9, -105, 21, -105, 11, -105, 31, -105, 35, -105, 37, -105,
	42, -105, 38, -105, 8, -105, 4, -105, 26, -105, 25, -105, 7, -105,
	26, 9, -43, 21, -43, 11, -43, 31, -43, 35, -43, 37, -43, 42,
	-43, 38, -43, 8, -43, 4, -43, 26, -43, 25, -43, 7, -43, 26,
	9, -44, 21, -44, 11, -44, 31, -44, 35, -44, 37, -44, 42, -44,
	38, -44, 8, -44, 4, -44, 26, -44, 25, -44, 7, -44, 26, 9,
	-41, 21, -41, 11, -41, 31, -41, 35, -41, 37, -41, 42, -41, 38,
	-41, 8, -41, 4, -41, 26, -41, 25, -41, 7, -41, 26, 9, -42,
	21, -42, 11, -42, 31, -42, 35, -42, 37, -42, 42, -42, 38, -42,
	8, -42, 4, -42, 26, -42, 25, -42, 7, -42, 26, 9, -40, 21,
	-40, 11, -40, 31, -40, 35, -40, 37, -40, 42, -40, 38, -40, 8,
	-40, 4, -40, 26, -40, 25, -40, 7, -40, 26, 9, -107, 21, -107,
	11, -107, 31, -107, 35, -107, 37, -107, 42, -107, 38, -107, 8, -107,
	4, -107, 26, -107, 25, -107, 7, -107, 26, 9, -39, 21, -39, 11,
	-39, 31, -39, 35, -39, 37, -39, 42, -39, 38, -39, 8, -39, 4,
	-39, 26, -39, 25, -39, 7, -39, 2, 9, 108, 4, 38, -110, 7,
	47, 4, 41, 110, 40, 111, 4, 42, 119, 4, 120, 4, 42, -10,
	4, -10, 28, 2, -14, 9, -14, 27, -14, 35, -14, 28, -14, 29,
	-14, 37, -14, 42, -14, 15, -14, 4, -14, 30, -14, 13, -14, 17,
	-14, 12, -14, 28, 2, -12, 9, -12, 27, -12, 35, -12, 28, -12,
	29, -12, 37, -12, 42, -12, 15, -12, 4, -12, 30, -12, 13, -12,
	17, -12, 12, -12, 2, 8, 132, 28, 2, -13, 9, -13, 27, -13,
	35, -13, 28, -13, 29, -13, 37, -13, 42, -13, 15, -13, 4, -13,
	30, -13, 13, -13, 17, -13, 12, -13, 28, 2, -15, 9, -15, 27,
	-15, 35, -15, 28, -15, 29, -15, 37, -15, 42, -15, 15, -15, 4,
	-15, 30, -15, 13, -15, 17, -15, 12, -15, 4, 42, -78, 4, -78,
	24, 27, -84, 35, -84, 28, -84, 29, -84, 37, -84, 42, -84, 15,
	126, 4, -84, 30, -84, 13, 127, 17, 128, 12, 129, 16, 27, -80,
	35, -80, 28, -80, 29, -80, 37, -80, 42, -80, 4, -80, 30, -80,
	16, 27, 95, 35, 96, 28, 121, 29, 97, 37, 98, 42, -82, 4,
	-82, 30, 122, 12, 6, -94, 33, 24, 23, 25, 35, 26, 22, 27,
	42, 29, 2, 42, 136, 20, 6, -33, 0, -33, 33, -33, 23, -33,
	35, -33, 19, -33, 22, -33, 24, -33, 42, -33, 18, -33, 2, 42,
	137, 36, 9, -49, 21, -49, 11, -49, 31, -49, 35, -49, 37, -49,
	42, -49, 38, -49, 15, -49, 16, -49, 8, -49, 4, -49, 26, -49,
	25, -49, 7, -49, 13, -49, 14, -49, 12, -49, 36, 9, -50, 21,
	-50, 11, -50, 31, -50, 35, -50, 37, -50, 42, -50, 38, -50, 15,
	-50, 16, -50, 8, -50, 4, -50, 26, -50, 25, -50, 7, -50, 13,
	-50, 14, -50, 12, -50, 6, 39, -53, 41, -53, 40, -53, 6, 39,
	-54, 41, -54, 40, -54, 6, 39, -112, 41, -112, 40, -112, 6, 39,
	138, 41, 110, 40, 111, 4, 9, -114, 35, 140, 2, 35, 142, 4,
	42, -91, 5, -91, 10, 21, -99, 31, -99, 42, -99, 26, -99, 25,
	-99, 14, 9, -103, 21, -103, 31, -103, 42, -103, 4, -103, 26, -103,
	25, -103, 12, 0, -8, 35, -8, 19, -8, 42, -8, 18, -8, 20,
	-8, 2, 8, 144, 2, 8, 145, 4, 42, -81, 4, -81, 4, 42,
	-9, 4, -9, 16, 27, -79, 35, -79, 28, -79, 29, -79, 37, -79,
	42, -79, 4, -79, 30, -79, 16, 27, -19, 35, -19, 28, -19, 29,
	-19, 37, -19, 42, -19, 4, -19, 30, -19, 16, 27, -17, 35, -17,
	28, -17, 29, -17, 37, -17, 42, -17, 4, -17, 30, -17, 16, 27,
	-18, 35, -18, 28, -18, 29, -18, 37, -18, 42, -18, 4, -18, 30,
	-18, 16, 27, -20, 35, -20, 28, -20, 29, -20, 37, -20, 42, -20,
	4, -20, 30, -20, 16, 27, -83, 35, -83, 28, -83, 29, -83, 37,
	-83, 42, -83, 4, -83, 30, -83, 16, 27, -11, 35, -11, 28, -11,
	29, -11, 37, -11, 42, -11, 4, -11, 30, -11, 8, 27, 95, 35,
	96, 29, 97, 37, 98, 12, 6, -96, 33, -96, 23, -96, 35, -96,
	22, -96, 42, -96, 2, 6, 146, 12, 6, -93, 33, 24, 23, 25,
	35, 26, 22, 27, 42, 29, 20, 6, -32, 0, -32, 33, -32, 23,
	-32, 35, -32, 19, -32, 22, -32, 24, -32, 42, -32, 18, -32, 20,
	6, -34, 0, -34, 33, -34, 23, -34, 35, -34, 19, -34, 22, -34,
	24, -34, 42, -34, 18, -34, 38, 9, -52, 21, -52, 11, -52, 31,
	-52, 35, -52, 37, -52, 42, -52, 38, -52, 15, -52, 16, -52, 8,
	-52, 4, -52, 26, -52, 25, -52, 10, -52, 7, -52, 13, -52, 14,
	-52, 12, -52, 6, 39, -111, 41, -111, 40, -111, 2, 9, -113, 2,
	9, 147, 2, 9, 148, 2, 2, 151, 2, 36, 152, 2, 36, 153,
	18, 0, -31, 33, -31, 23, -31, 35, -31, 19, -31, 22, -31, 24,
	-31, 42, -31, 18, -31, 10, 21, -60, 31, -60, 42, -60, 26, -60,
	25, -60, 10, 21, -62, 31, -62, 42, -62, 26, -62, 25, -62, 4,
	42, -77, 4, -77, 12, 6, -95, 33, -95, 23, -95, 35, -95, 22,
	-95, 42, -95, 2, 9, 155, 2, 9, 156, 2, 9, 157, 4, 42,
	-21, 4, -21, 4, 42, -22, 4, -22, 28, 2, -16, 9, -16, 27,
	-16, 35, -16, 28, -16, 29, -16, 37, -16, 42, -16, 15, -16, 4,
	-16, 30, -16, 13, -16, 17, -16, 12, -16,
}

var _goto = []int32{
	158, 165, 165, 166, 165, 165, 165, 165, 165, 165, 165, 165, 177, 165,
	184, 195, 165, 165, 165, 165, 165, 165, 165, 214, 221, 226, 165, 165,
	165, 165, 165, 165, 165, 165, 165, 245, 165, 165, 165, 165, 260, 265,
	165, 165, 165, 165, 284, 165, 165, 165, 165, 303, 165, 318, 165, 323,
	165, 165, 165, 334, 165, 165, 337, 165, 165, 165, 350, 165, 165, 165,
	165, 165, 165, 365, 165, 165, 165, 165, 376, 165, 165, 395, 165, 165,
	165, 165, 165, 165, 165, 165, 165, 410, 415, 165, 165, 165, 165, 165,
	165, 165, 165, 420, 165, 425, 436, 165, 165, 165, 165, 165, 165, 165,
	165, 451, 454, 165, 165, 165, 165, 165, 457, 165, 165, 165, 165, 165,
	165, 165, 165, 165, 165, 165, 468, 165, 165, 473, 165, 165, 165, 165,
	165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 484, 165, 165,
	165, 165, 165, 165, 6, 34, 3, 35, 4, 1, 5, 0, 10, 12,
	8, 3, 9, 2, 10, 36, 11, 37, 12, 6, 12, 8, 3, 9,
	2, 16, 10, 40, 19, 5, 20, 4, 21, 38, 22, 39, 23, 18,
	19, 30, 17, 31, 14, 32, 13, 33, 45, 34, 46, 35, 18, 36,
	15, 37, 16, 38, 6, 40, 19, 5, 20, 4, 60, 4, 20, 58,
	53, 59, 18, 54, 42, 57, 48, 27, 49, 26, 50, 21, 51, 22,
	52, 25, 53, 23, 54, 55, 55, 14, 19, 30, 17, 31, 14, 32,
	13, 61, 18, 36, 15, 37, 16, 38, 4, 47, 64, 48, 65, 18,
	54, 42, 57, 48, 27, 49, 26, 50, 21, 66, 22, 52, 25, 53,
	23, 54, 55, 55, 18, 54, 42, 57, 48, 27, 49, 26, 50, 21,
	90, 22, 52, 25, 53, 23, 54, 55, 55, 14, 29, 71, 51, 72,
	52, 73, 30, 74, 33, 75, 32, 76, 31, 77, 4, 24, 88, 56,
	89, 10, 57, 48, 27, 49, 26, 50, 25, 53, 23, 82, 2, 20,
	80, 12, 41, 93, 9, 99, 6, 100, 8, 101, 7, 102, 42, 103,
	14, 29, 71, 51, 105, 52, 73, 30, 74, 33, 75, 32, 76, 31,
	77, 10, 29, 117, 30, 74, 33, 75, 32, 76, 31, 77, 18, 54,
	42, 57, 48, 27, 49, 26, 50, 21, 107, 22, 52, 25, 53, 23,
	54, 55, 55, 14, 57, 48, 27, 49, 26, 50, 22, 118, 25, 53,
	23, 54, 55, 55, 4, 57, 48, 27, 109, 4, 28, 112, 58, 113,
	4, 10, 130, 44, 131, 10, 9, 99, 11, 123, 43, 124, 8, 101,
	7, 125, 14, 19, 30, 17, 31, 14, 133, 49, 134, 50, 135, 18,
	36, 16, 38, 2, 28, 139, 2, 59, 141, 10, 9, 99, 6, 149,
	8, 101, 7, 102, 42, 103, 4, 9, 99, 8, 143, 10, 19, 30,
	17, 31, 14, 150, 18, 36, 16, 38, 4, 9, 99, 8, 154,
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
		return p.on_lexer_term__tok(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 48:
		return p.on_lexer_term__char_class_expr(
			_cast[_i0.CharClassExpr](p._stack.Peek(0).Sym),
		)
	case 49:
		return p.on_lexer_term__expr(
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[*_i0.LexerExpr](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 50:
		return p.on_char_class_expr__binary(
			_cast[_i0.CharClassExpr](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[_i0.CharClassExpr](p._stack.Peek(0).Sym),
		)
	case 51:
		return p.on_char_class_expr__char_class(
			_cast[*_i0.CharClass](p._stack.Peek(0).Sym),
		)
	case 52:
		return p.on_char_class(
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[[]Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 53:
		return p.on_char_class_item(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 54:
		return p.on_char_class_item(
			_cast[Token](p._stack.Peek(0).Sym),
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
		return p.on_action(
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 59:
		return p.on_action_discard(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 60:
		return p.on_action_push_mode(
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 61:
		return p.on_action_pop_mode(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 62:
		return p.on_action_emit(
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 63: // ZeroOrMore
		return _cast[[]Token](p._stack.Peek(0).Sym)
	case 64: // ZeroOrMore
		{
			var zero []Token
			return zero
		}
	case 65:
		{ // OneOrMoreF
			l := _cast[[]Token](p._stack.Peek(1).Sym)
			e := _cast[Token](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 66:
		{ // OneOrMoreF
			var l []Token
			e := _cast[Token](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 67: // ZeroOrMore
		return _cast[[][]_i0.Statement](p._stack.Peek(0).Sym)
	case 68: // ZeroOrMore
		{
			var zero [][]_i0.Statement
			return zero
		}
	case 69: // OneOrMore
		return append(
			_cast[[][]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 70: // OneOrMore
		return [][]_i0.Statement{
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		}
	case 71: // ZeroOrMore
		return _cast[[]_i0.Statement](p._stack.Peek(0).Sym)
	case 72: // ZeroOrMore
		{
			var zero []_i0.Statement
			return zero
		}
	case 73: // OneOrMore
		return append(
			_cast[[]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 74: // OneOrMore
		return []_i0.Statement{
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		}
	case 75: // ZeroOrOne
		return _cast[Token](p._stack.Peek(0).Sym)
	case 76: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	case 77: // List
		return append(
			_cast[[]*_i0.ParserProd](p._stack.Peek(2).Sym),
			_cast[*_i0.ParserProd](p._stack.Peek(0).Sym),
		)
	case 78: // List
		return []*_i0.ParserProd{
			_cast[*_i0.ParserProd](p._stack.Peek(0).Sym),
		}
	case 79: // OneOrMore
		return append(
			_cast[[]*_i0.ParserTerm](p._stack.Peek(1).Sym),
			_cast[*_i0.ParserTerm](p._stack.Peek(0).Sym),
		)
	case 80: // OneOrMore
		return []*_i0.ParserTerm{
			_cast[*_i0.ParserTerm](p._stack.Peek(0).Sym),
		}
	case 81: // ZeroOrOne
		return _cast[*_i0.ProdQualifier](p._stack.Peek(0).Sym)
	case 82: // ZeroOrOne
		{
			var zero *_i0.ProdQualifier
			return zero
		}
	case 83: // ZeroOrOne
		return _cast[_i0.ParserTermType](p._stack.Peek(0).Sym)
	case 84: // ZeroOrOne
		{
			var zero _i0.ParserTermType
			return zero
		}
	case 85: // ZeroOrMore
		return _cast[[]_i0.Statement](p._stack.Peek(0).Sym)
	case 86: // ZeroOrMore
		{
			var zero []_i0.Statement
			return zero
		}
	case 87:
		{ // OneOrMoreF
			l := _cast[[]_i0.Statement](p._stack.Peek(1).Sym)
			e := _cast[_i0.Statement](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 88:
		{ // OneOrMoreF
			var l []_i0.Statement
			e := _cast[_i0.Statement](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 89: // ZeroOrMore
		return _cast[[]Token](p._stack.Peek(0).Sym)
	case 90: // ZeroOrMore
		{
			var zero []Token
			return zero
		}
	case 91: // OneOrMore
		return append(
			_cast[[]Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 92: // OneOrMore
		return []Token{
			_cast[Token](p._stack.Peek(0).Sym),
		}
	case 93: // ZeroOrMore
		return _cast[[]_i0.Statement](p._stack.Peek(0).Sym)
	case 94: // ZeroOrMore
		{
			var zero []_i0.Statement
			return zero
		}
	case 95: // OneOrMore
		return append(
			_cast[[]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 96: // OneOrMore
		return []_i0.Statement{
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		}
	case 97: // ZeroOrMore
		return _cast[[]_i0.Action](p._stack.Peek(0).Sym)
	case 98: // ZeroOrMore
		{
			var zero []_i0.Action
			return zero
		}
	case 99: // OneOrMore
		return append(
			_cast[[]_i0.Action](p._stack.Peek(1).Sym),
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 100: // OneOrMore
		return []_i0.Action{
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		}
	case 101: // OneOrMore
		return append(
			_cast[[]*_i0.ExternalName](p._stack.Peek(1).Sym),
			_cast[*_i0.ExternalName](p._stack.Peek(0).Sym),
		)
	case 102: // OneOrMore
		return []*_i0.ExternalName{
			_cast[*_i0.ExternalName](p._stack.Peek(0).Sym),
		}
	case 103: // List
		return append(
			_cast[[]*_i0.LexerFactor](p._stack.Peek(2).Sym),
			_cast[*_i0.LexerFactor](p._stack.Peek(0).Sym),
		)
	case 104: // List
		return []*_i0.LexerFactor{
			_cast[*_i0.LexerFactor](p._stack.Peek(0).Sym),
		}
	case 105: // OneOrMore
		return append(
			_cast[[]*_i0.LexerTermCard](p._stack.Peek(1).Sym),
			_cast[*_i0.LexerTermCard](p._stack.Peek(0).Sym),
		)
	case 106: // OneOrMore
		return []*_i0.LexerTermCard{
			_cast[*_i0.LexerTermCard](p._stack.Peek(0).Sym),
		}
	case 107: // ZeroOrOne
		return _cast[_i0.Card](p._stack.Peek(0).Sym)
	case 108: // ZeroOrOne
		{
			var zero _i0.Card
			return zero
		}
	case 109: // ZeroOrOne
		return _cast[Token](p._stack.Peek(0).Sym)
	case 110: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	case 111: // OneOrMore
		return append(
			_cast[[]Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 112: // OneOrMore
		return []Token{
			_cast[Token](p._stack.Peek(0).Sym),
		}
	case 113: // ZeroOrOne
		return _cast[Token](p._stack.Peek(0).Sym)
	case 114: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	default:
		panic("unreachable")
	}
}
