import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = [
    "[({(<(())[]>[[{[]{<()<>>",
    "[(()[<>])]({[<{<<[]>>(",
    "{([(<{}[<>[]}>{[]{[(<()>",
    "(((({<>}<{<{<>}{[]{[]{}",
    "[[<[([]))<([[{}[[()]]]",
    "[{[{({}]{}}([{[{{{}}([]",
    "{<[[]]>}<{[{[{[]{()[[[]",
    "[<(<(<(<{}))><([]([]()",
    "<{([([[(<>()){}]>(<<{{",
    "<{([{{}}[<[[[<>{}]]]>[]]",
]
ACTUAL_INPUT = 'input/day_10.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(10, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '26397'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '392097'


@pytest.mark.parametrize('solution', [
    GoSolution(10, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '288957'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == ''
