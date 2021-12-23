import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = 'input/day_22_example.txt'
OFFICIAL_EXAMPLE_2 = 'input/day_22_example_2.txt'
ACTUAL_INPUT = 'input/day_22.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(22, 1),
    GoSolution(22, "1d"),
], ids=lambda s: s.name)
class TestPart1(object):

    @pytest.mark.parametrize('lines, expected', (
        (["on x=0..1,y=0..1,z=0..1"], "8"),
        ([
             "on x=0..1,y=0..1,z=0..1",
             "on x=10..11,y=10..11,z=10..11"
         ], "16"),
        ([
             "on x=0..1,y=0..1,z=0..1",
             "off x=0..1,y=0..1,z=0..1"
         ], "0"),
        ([
            "on x=0..1,y=0..1,z=0..1",
            "off x=0..1,y=0..1,z=1..2"
         ], "4"),
        ([
             "on x=0..1,y=0..1,z=0..1",
             "off x=1..10,y=1..10,z=1..10"
         ], "7"),
        ([
             "on x=0..2,y=0..2,z=0..2",
             "off x=0..1,y=0..1,z=0..1"
         ], "19"),
        ([
             "on x=0..2,y=0..2,z=0..2",
             "on x=0..1,y=0..1,z=0..5"
         ], "39"),
    ))
    def test_minimal_examples(self, solution, lines, expected):
        assert solution.run_lines(lines) == expected

    def test_official_example(self, solution):
        assert solution.run_file(OFFICIAL_EXAMPLE) == '590784'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '606484'


@pytest.mark.parametrize('solution', [
    GoSolution(22, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_file(OFFICIAL_EXAMPLE_2) == '2758514936282235'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '1162571910364852'
