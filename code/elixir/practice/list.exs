defmodule MyList do
  def len([]),          do: 0
  def len([_ | tail]),  do: 1 + len(tail)
  def mapsum([], _),    do: 0
  def mapsum([ head | tail ], func) do
    Enum.reduce(tail, head, &(func.(&1) + &2))
  end
end

IO.puts MyList.len([1,2,3])

IO.puts MyList.mapsum [1,2,3], &(&1 * &1)
IO.puts MyList.mapsum [], &(&1 * &1)
