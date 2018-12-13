#!/usr/bin/env python
# -*- coding: utf-8 -*-

import matplotlib
import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
import sklearn

oecd_bli = pd.read_csv("oecd_bli_2015.csv", thousands=',')
gdp_per_capita = pd.read_csv("gdp_per_capita.csv",thousands=',',delimiter='\t', encoding='latin1', na_values="n/a")

def prepare_country_stats(oecd_bli, gdp_per_capita):
    full_country_stats = pd.merge(left=oecd_bli, right=gdp_per_capita, left_index=True, right_index=True)
    full_country_stats.sort_values(by="GDP per capita", inplace="True")
    full_country_stats[["GDP per capita", 'Life satisfaction']].loc["United States"]
    remove_indices = [0, 1, 6, 8, 33, 34, 35]
    keep_indices = list(set(range(36)) - set(remove_indices))
    sample_data = full_country_stats[["GDP per capita", 'Life satisfaction']].iloc[keep_indices]
    sample_data.plot(kind='scatter', x="GDP per capita", y='Life satisfaction', figsize=(5,3))
    return sample_data

country_stats = prepare_country_stats(oecd_bli, gdp_per_capita)
X = np.c_[country_stats["GDP per capita"]]
y = np.c_[country_stats["Life satisfaction"]]

country_stats.plot(kind='scatter', x="GDP per capita", y='Life satisfaction')
plt.show()

lin_reg_model = sklearn.linear_model.LinearRegression()

lin_reg_model.fit(X, y)

X_new = [[22587]]
print(lin_reg_model.predict(X_new))
