import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
# read data into dataframe
df = pd.read_csv('ViralSpread.exe_assign_no_treatment.out.csv')
# plot data
plt.plot(df['Time (day)'], df['Number of normal cells'], label='Uninfected')
plt.plot(df['Time (day)'], df['Number of target cells'], label='Infected')
plt.plot(df['Time (day)'], df['Number of infectious cells'], label='Infectious')
plt.plot(df['Time (day)'], df['Number of dead cells'], label='Dead')
# setup plot
plt.rcParams["figure.figsize"] = (8,8)
plt.rcParams.update({'font.size': 15})
plt.title('Number of Cells vs. Time')
plt.xlabel('Time (days)')
plt.ylabel('Number of cells')
plt.legend()
plt.savefig('ViralSpread.exe_assign_no_treatment.out.jpg',bbox_inches='tight', dpi=500)