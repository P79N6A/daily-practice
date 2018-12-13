defmodule Main do
  def fizzbuzz(n) when n > 0, do: _downto(n, [])

  defp _downto(0, result), do: result
  defp _downto(n, result) do
    next =
      cond do
        rem(n, 3) == 0 and rem(n, 5) == 0 ->
          "FizzBuzz"
        rem(n, 3) == 0 ->
          "Fizz"
        rem(n, 5) == 0 ->
          "Buzz"
        true ->
          n
      end
    _downto(n-1, [ next | result])
  end
end

IO.inspect Main.fizzbuzz(16)
