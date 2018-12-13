defmodule WeatherHistory do
  def for_location([], _target_loc), do: []

  def for_location([ head = [ _, target_loc, _, _ ] | tail], target_loc) do
    [ head | for_location(tail, target_loc) ]
  end

  def for_location([ _ | tail ], target_loc), do: for_location(tail, target_loc)

  def test_data do
    [
      [133645354, 26, 16, 0.125],
      [133645353, 26, 16, 0.125],
      [133645355, 27, 16, 0.125],
      [133645357, 26, 16, 0.125],
      [133645358, 27, 16, 0.125],
      [133645359, 28, 16, 0.125],
      [133645350, 26, 16, 0.125],
      [133645351, 28, 16, 0.125],
    ]
  end
end



IO.inspect  WeatherHistory.for_location(WeatherHistory.test_data, 27)
