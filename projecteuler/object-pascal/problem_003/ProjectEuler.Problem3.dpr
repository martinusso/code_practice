program ProjectEuler.Problem3;

{$APPTYPE CONSOLE}

uses
  SysUtils;

var
  Number: Int64;
  I: Int64;
begin
  Number := 600851475143;
  I := 2;
  while I * I < Number do
  begin
    while Number mod I = 0 do
      Number := Round(Number / I);

    Inc(I);
  end;
  Writeln(IntToStr(Number));
end.
