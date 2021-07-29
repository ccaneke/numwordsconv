package numwordsconv

var testCases []struct {
	in   int
	want string
} = []struct {
	in   int
	want string
}{{1, "one"},
	{56, "fifty six"},
	{126, "one hundred and twenty six"},
	{1012, "one thousand and twelve"},
	{1026, "one thousand and twenty six"},
	{1526, "one thousand five hundred and twenty six"},
	{9999, "nine thousand nine hundred and ninety nine"}}
