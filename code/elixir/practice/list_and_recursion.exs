defmodule ListAndRecursion do
  def sum([], total),               do: total
  def sum([ head | tail ], total),  do: sum(tail, head+total)

  # practice 10-5
  def filter([head|tail], f) do
    if f.(head) do
      [head | filter(tail, f)]
    else
      filter(tail, f)
    end
  end
  def filter([], _) do
    []
  end
  # practice 10-6
  def flatten([head|tail]) do
    if is_list(head) do
      flatten(head) ++ flatten(tail)
    else
      [head| flatten(tail)]
    end
  end
  def flatten([]) do
    []
  end
end

IO.puts ListAndRecursion.sum([1,2,3,4,5,6], 0)
IO.inspect ListAndRecursion.filter([1,2,3,4,5], fn v -> v > 3 end)
IO.inspect ListAndRecursion.flatten([1, [3, 4, [2]], 7, [[[6]]]])
