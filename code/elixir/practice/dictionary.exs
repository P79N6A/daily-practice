defmodule Sum do
  def values(dict) do
    dict |> Map.values |> Enum.sum
  end
end

defmodule Data do
  def people do
    [
      %{ name: "Grumpy",      height: 1.24 },
      %{ name: "Dave",        height: 1.88 },
      %{ name: "Dopey",       height: 1.32 },
      %{ name: "Shaquille",   height: 2.16 },
      %{ name: "Sneezy",      height: 1.28 }
    ]
  end
end

[ one: 1, two: 2, three: 3 ] |> Enum.into(Map.new) |> Sum.values |> IO.puts
%{ four: 4, five: 5, six: 6 } |> Enum.into(Map.new) |> Sum.values |> IO.puts
[ point: 2, point: 3, name: "Dave"] |> Keyword.get_values(:point) |> IO.inspect

for person = %{ height: height} <- Data.people,
  height > 1.5,
  do: IO.inspect person

defmodule HotelRoom do
  def book(%{name: name, height: height})
  when height > 1.5 do
    IO.puts "Need extra long bed for #{name}"
  end

  def book(%{name: name, height: height})
  when height < 1.3 do
    IO.puts "Need low shower controls for #{name}"
  end

  def book(person) do
    IO.puts "Need regular bed for #{person.name}"
  end
end

Data.people |> Enum.each(&HotelRoom.book/1)
