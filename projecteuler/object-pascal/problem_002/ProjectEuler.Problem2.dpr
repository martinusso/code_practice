program ProjectEuler.Problem2;

{$APPTYPE CONSOLE}

uses
  SysUtils;

const
  MAXIMUM = 4000000;
var
  fibonacciNumber: Integer;
  firstPrevious: Integer;
  secondPrevious: Integer;
  value: Integer;
begin
  fibonacciNumber := 1;
  firstPrevious := 1;
  secondPrevious := 0;
  while fibonacciNumber <= MAXIMUM do
  begin
    fibonacciNumber := firstPrevious + secondPrevious;
    secondPrevious := firstPrevious;
    firstPrevious := fibonacciNumber;
    value := value + fibonacciNumber;
  end;
  Writeln(IntToStr(value));
end.
