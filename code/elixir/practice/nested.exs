defmodule Main do

  def nested do
    %{
      buttercup: %{
        actor: %{
          first: "Robin",
          last: "Wright"
        },
        role: "princess"
      },
      westley: %{
        actor: %{
          first: "Carey",
          last: "Ewes"
        },
        role: "farm boy"
      }
    }
  end

  defmodule Customer do
    defstruct name: "", company: ""
  end

  defmodule BugReport do
    defstruct owner: %{}, details: "", severity: 1
  end

  def main do
    report = %BugReport{ owner: %Customer{name: "Dave", company: "Pragmatic"}, details: "broken" }
    update_in(report.owner.name, &("Mr. " <> &1))
    IO.inspect report
  end

end

Main.main

nested_new = put_in(Main.nested, [:westley, :actor, :last], "Elwes")
IO.inspect get_in(nested_new, [:westley, :actor, :last])
