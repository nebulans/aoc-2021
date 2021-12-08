import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = ["16,1,2,0,4,2,7,1,2,14"]
ACTUAL_INPUT = 'input/day_07.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(7, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '37'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '343605'


@pytest.mark.parametrize('solution', [
    GoSolution(7, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '168'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '96744904'
