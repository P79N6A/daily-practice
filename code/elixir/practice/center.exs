defmodule Main do

  def center(words)
  when is_list(words) do
    max = String.length(Enum.max_by(words, fn(x) -> String.length(x) end))
    center_internal(words, max)
  end

  defp center_internal([head|tail], max)
  when is_bitstring(head) do
    len = String.length(head)
    padding = round(:math.floor(len + (max - len) / 2))
    IO.puts String.pad_leading(head, padding)
    center_internal(tail, max)
  end

  defp center_internal([], _) do
    nil
  end
end

Main.center(["cat", "zebra", "elephant"])
