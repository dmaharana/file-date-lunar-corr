# File Date Lunar Correlation Specification

This project implements a CLI tool to recursively scan files, extract their creation dates, and analyze their correlation with lunar cycles.

## Objective
The primary goal is to test whether file creation dates cluster around particular lunar phases (e.g., more files created near a full moon or new moon) or if the distribution is statistically random.

## General Approach

1.  **Date Extraction**: Recursively traverse subdirectories and retrieve the creation date (birthtime) for each file.
2.  **Lunar Phase Mapping**: Map each date to a lunar phase (0.0 - 1.0, where 0 is New Moon and 0.5 is Full Moon).
3.  **Phase Binning**: Group the continuous phase values into 8 distinct segments:
    - New Moon
    - Waxing Crescent
    - First Quarter
    - Waxing Gibbous
    - Full Moon
    - Waning Gibbous
    - Last Quarter
    - Waning Crescent
4.  **Statistical Analysis**: Perform a **Chi-squared goodness-of-fit test** to check if the observed counts per phase differ significantly from a uniform (random) distribution.

## Statistical Interpretation

-   **p > 0.05**: No significant correlation found. The distribution of file creation dates across lunar phases appears random.
-   **p ≤ 0.05**: The distribution is statistically non-uniform, suggesting a possible pattern or correlation worth investigating.

## Reference Implementation (Python)

The following Python snippet illustrates the logic used for calculation and analysis:

```python
import pandas as pd
import ephem
from scipy.stats import chisquare

# Load and clean data
df = pd.read_csv("file_dates.csv")
df["Created Date"] = pd.to_datetime(df["Created Date"], errors="coerce")
df = df.dropna(subset=["Created Date"])

def lunar_phase(dt):
    moon = ephem.Moon(dt.strftime("%Y/%m/%d"))
    return moon.phase / 100  # Returns 0-1 illumination %

df["lunar_phase"] = df["Created Date"].apply(lunar_phase)

# Bin into 8 phases
bins = [0, .125, .25, .375, .5, .625, .75, .875, 1.0]
labels = ["New","Wax-Crescent","First-Qtr","Wax-Gibbous",
          "Full","Wan-Gibbous","Last-Qtr","Wan-Crescent"]
df["phase_name"] = pd.cut(df["lunar_phase"], bins=bins, labels=labels)

counts = df["phase_name"].value_counts().sort_index()
print(counts)

# Chi-squared test
chi2, p = chisquare(counts.values)
print(f"\nChi² = {chi2:.2f}, p-value = {p:.4f}")
print("Likely random." if p > 0.05 else "Statistically non-uniform — possible pattern!")
```

## Important Considerations
Even a statistically significant result may reflect external factors such as a standard work schedule (weekdays vs. weekends) rather than the influence of the moon. Since lunar cycles (~29.5 days) do not align perfectly with the 7-day calendar week, further analysis controlling for the day of the week is recommended to rule out such biases.
