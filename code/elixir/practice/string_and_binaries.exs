defmodule Main do

  def capitalize_sentences(sentence) do
    pat = ". "
    String.split(sentence, pat)
      |> Enum.map(fn s -> String.capitalize(s) end)
      |> Enum.join(pat)
  end
end

IO.puts Main.capitalize_sentences("oh. a Dog. woof. ")
