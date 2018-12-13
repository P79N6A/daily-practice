defmodule Main do
  def main do
    first8 = Enum.into 1..8, []
    IO.inspect for x <- first8, y <- first8, x >= y, rem(x*y, 10) == 0, do: { x, y }
  end
end

Main.main
