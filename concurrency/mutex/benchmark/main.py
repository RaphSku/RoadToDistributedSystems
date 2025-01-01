import re

import pandas as pd
import plotly.graph_objects as go


def convert_to_ms(value, unit):
    if unit == "µs":
        return value * 0.001
    elif unit == "ns":
        return value * 0.000001
    elif unit == "s":
        return value * 1000
    return value


def main():
    data = {"targets": [], "time_ms_parallel": [], "time_ms_sequential": []}
    with open("benchmark.log", "r") as file:
        lines = file.readlines()

    samples_pattern = re.compile(r"Number of samples for each target:\s*(\d+)")
    target_pattern = re.compile(r"Target:\s*(\d+)")
    parallel_pattern = re.compile(r"Elapsed Time \(parallel\):\s*([\d.]+)(\w+)")
    sequential_pattern = re.compile(r"Elapsed Time \(sequential\):\s*([\d.]+)(\w+)")

    samples_averaged_over = None
    for line in lines:
        samples_match = samples_pattern.search(line)
        if samples_match:
            samples_averaged_over = int(samples_match.group(1))

        target_match = target_pattern.search(line)
        if target_match:
            current_target = int(target_match.group(1))
            data["targets"].append(current_target)

        parallel_match = parallel_pattern.search(line)
        if parallel_match:
            elapsed_parallel_value = float(parallel_match.group(1))
            unit = parallel_match.group(2)
            data["time_ms_parallel"].append(
                round(convert_to_ms(elapsed_parallel_value, unit), 3)
            )

        sequential_match = sequential_pattern.search(line)
        if sequential_match:
            elapsed_sequential_value = float(sequential_match.group(1))
            unit = sequential_match.group(2)
            data["time_ms_sequential"].append(
                round(convert_to_ms(elapsed_sequential_value, unit), 3)
            )

    trace_parallel = go.Scatter(
        x=data["targets"],
        y=data["time_ms_parallel"],
        mode="lines+markers",
        name="Elapsed Time (Parallel)",
        line={
            "color": "blue",
        },
        marker={
            "size": 10,
        },
    )

    trace_sequential = go.Scatter(
        x=data["targets"],
        y=data["time_ms_sequential"],
        mode="lines+markers",
        name="Elapsed Time (Sequential)",
        line={
            "color": "red",
        },
        marker={
            "size": 10,
        },
    )

    layout = go.Layout(
        title=f"Elapsed Time for Different Targets, Samples per Target: {samples_averaged_over}",
        title_font={
            "size": 30,
        },
        xaxis={
            "title": "Targets (log scaled)",
            "type": "log",
            "title_font": {
                "size": 24,
            },
            "tickfont": {
                "size": 18,
            },
        },
        yaxis={
            "title": "Elapsed Time (ms)",
            "title_font": {
                "size": 24,
            },
            "tickfont": {
                "size": 18,
            },
        },
        showlegend=True,
        legend={
            "font": {
                "size": 16,
            },
        },
    )

    fig = go.Figure(data=[trace_parallel, trace_sequential], layout=layout)
    fig.show()


if __name__ == "__main__":
    main()
