program ProjectEuler.Problem2;

{$APPTYPE CONSOLE}

uses
  SysUtils;

const
  MAXIMUM = 4000000;
var
  F, PreviousTerm, LastTerm: Integer;
  Sum: Integer;
begin
  PreviousTerm := 0;
  LastTerm := 1;
  while LastTerm <= MAXIMUM do
  begin
    F := PreviousTerm + LastTerm;
    PreviousTerm := LastTerm;
    LastTerm := F;
    if LastTerm mod 2 = 0 then
      Sum := Sum + LastTerm;
  end;
  Writeln(IntToStr(Sum));
  Readln;
end.
