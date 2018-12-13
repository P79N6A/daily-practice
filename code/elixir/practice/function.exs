defmodule Fns do
  def sum(a, b) do
    a + b
  end
end

handle_open = fn
  {:ok, file} -> "First line: #{IO.read(file, :line)}"
  {_, error} -> "Error: #{:file.format_error(error)}"
end

IO.puts Fns.sum(1, 2)
IO.puts handle_open.(File.open("/etc/hosts"))

fizzbuzz = fn
  (0, 0, _) -> "fizzbuzz"
  (0, _, _) -> "fizz"
  (_, 0, _) -> "buzz"
  (_, _, a) -> a
end

IO.puts fizzbuzz.(0, 0, 3)
IO.puts fizzbuzz.(0, 2, 3)
IO.puts fizzbuzz.(1, 0, 3)
IO.puts fizzbuzz.(1, 2, 3)

prefix = fn s1 -> (fn s2 -> s1 ++ s2 end) end
IO.puts prefix.('hello').(' ,world')

IO.inspect Enum.map [1,2,3,4], &(&1+2)
