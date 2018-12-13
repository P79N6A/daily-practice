defmodule Main do
  def main do
    1..5
    |> Enum.to_list
    |> Enum.map(&(&1*&1))
    |> Enum.with_index
    |> Enum.map(fn {value, index} -> value - index end)
    |> IO.inspect
  end

  def fibonacci(n) do
    Stream.unfold({0,1}, fn {acc, next} -> {acc, {next, acc+next}} end)
      |> Enum.at(n-1)
  end
end

Main.main
l = Enum.to_list(Stream.map [1,2,3,4,5,6],  &(&1+1))
IO.inspect l

IO.puts Main.fibonacci(5)
