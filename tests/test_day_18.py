import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = [
    "[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
    "[[[5,[2,8]],4],[5,[[9,9],0]]]",
    "[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
    "[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
    "[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
    "[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
    "[[[[5,4],[7,7]],8],[[8,3],8]]",
    "[[9,3],[[9,9],[6,[4,9]]]]",
    "[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
    "[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
]
ACTUAL_INPUT = 'input/day_18.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(18, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    @pytest.mark.parametrize('number, magnitude', [
        ("[1,1]", "5"),
        ("[2,2]", "10"),
        ("[[1,1],1]", "17"),
        ("[1,[1,1]]", "13"),
        ("[[1,2],[[3,4],5]]", "143"),
        ("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", "1384"),
        ("[[[[1,1],[2,2]],[3,3]],[4,4]]", "445"),
        ("[[[[3,0],[5,3]],[4,4]],[5,5]]", "791"),
        ("[[[[5,0],[7,4]],[5,5]],[6,6]]", "1137"),
        ("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", "3488"),
    ])
    def test_magnitude_examples(self, solution, number, magnitude):
        assert solution.run_lines([number]) == magnitude

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '4140'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '2501'


@pytest.mark.parametrize('solution', [
    GoSolution(18, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '3993'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '4935'
