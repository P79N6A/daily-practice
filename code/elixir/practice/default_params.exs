defmodule Example do
  def func(p1) when is_atom(p1) do
    IO.puts p1
  end

  def func(p1, p2 \\ 2, p3 \\ 3) do
    IO.inspect [p1, p2, p3]
  end
end

Example.func(1, 4)
Example.func(:a)
