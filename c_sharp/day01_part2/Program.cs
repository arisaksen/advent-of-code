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

var counts2 = new Dictionary<int, int>();

foreach (var line in lines)
{
    var parts = line.Split(" ", StringSplitOptions.RemoveEmptyEntries);
    if (parts.Length < 2)
    {
        Console.WriteLine($"Skipping invalid line: {line}");
        continue;
    }

    if (int.TryParse(parts[0], out var number1))
    {
        list1.Add(number1);
    }

    if (int.TryParse(parts[1], out var number2))
    {
        list2.Add(number2);
        if (counts2.ContainsKey(number2))
        {
            counts2[number2]++;
        }
        else
        {
            counts2[number2] = 1;
        }
    }
}

if (list1.Count != list2.Count)
{
    Console.WriteLine("Error: Unequal number of elements in the two lists.");
    return;
}

var total = 0;
for (var i = 0; i < list1.Count; i++)
{
    if (counts2.ContainsKey(list1[i]))
    {
        total += list1[i] * counts2[list1[i]];
    }
}

stopwatch.Stop();

Console.WriteLine($"Total Difference: {total}");
Console.WriteLine($"Execution Time: {stopwatch.ElapsedMilliseconds} ms");

var ticks = stopwatch.ElapsedTicks;
var microseconds = (double)ticks / Stopwatch.Frequency * 1000000;
Console.WriteLine($"Execution Time: {microseconds} µs");