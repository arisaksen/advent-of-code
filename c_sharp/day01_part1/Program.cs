using System.Diagnostics;

if (args.Length == 0)
{
    Console.WriteLine("Error: No input file provided.");
    return;
}

var filePath = args[0];
if (!File.Exists(filePath))
{
    Console.WriteLine($"Error: File '{filePath}' does not exist.");
    return;
}

// Start measuring time
var stopwatch = Stopwatch.StartNew();

var lines = File.ReadAllLines(filePath);
var list1 = new List<int>();
var list2 = new List<int>();

foreach (var line in lines)
{
    var parts = line.Split(" ", StringSplitOptions.RemoveEmptyEntries);
    if (parts.Length < 2)
    {
        Console.WriteLine($"Skipping invalid line: {line}");
        continue;
    }

    if (int.TryParse(parts[0], out var number1))
        list1.Add(number1);

    if (int.TryParse(parts[1], out var number2))
        list2.Add(number2);
}

if (list1.Count != list2.Count)
{
    Console.WriteLine("Error: Unequal number of elements in the two lists.");
    return;
}

list1.Sort();
list2.Sort();

var totalDiff = 0;
for (var i = 0; i < list1.Count; i++)
{
    totalDiff += Math.Abs(list1[i] - list2[i]);
}

stopwatch.Stop();

Console.WriteLine($"Total Difference: {totalDiff}");
Console.WriteLine($"Execution Time: {stopwatch.ElapsedMilliseconds} ms");

var ticks = stopwatch.ElapsedTicks;
var microseconds = (double)ticks / Stopwatch.Frequency * 1000000;
Console.WriteLine($"Execution Time: {microseconds} µs");