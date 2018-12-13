defmodule Main do
  defmodule User do
    defstruct name: "", paid: false, over_18: true

    def fetch(tup, key), do: {key, tup}
  end

  defimpl String.Chars, for: User do
    def to_string(user) do
      "User: #{user.name}, paid: #{user.paid}, over_18: #{user.over_18}"
    end
  end

  def main do
    user = %User{ name: "John", over_18: true, paid: false }
    IO.puts user
    IO.puts user.name
  end
end

Main.main
