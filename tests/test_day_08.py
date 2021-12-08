import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = 'input/day_08_example.txt'
ACTUAL_INPUT = 'input/day_08.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(8, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_official_example(self, solution):
        assert solution.run_file(OFFICIAL_EXAMPLE) == '26'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '321'


@pytest.mark.parametrize('solution', [
    GoSolution(8, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_file(OFFICIAL_EXAMPLE) == '61229'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '1028926'