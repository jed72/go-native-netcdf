# NetCDF File Information

## Basic Information
- **Filename**: pypsa-usa-caiso-2025-demo.nc
- **Path**: /Users/jed/Documents/pypsa-usa-caiso-2025-demo.nc
- **Size**: 8.03 MB

## Dimensions
- **snapshots**: 56
- **investment_periods**: 1
- **stores_i**: 12
- **stores_t_p_i**: 12
- **stores_t_e_i**: 12
- **stores_t_mu_lower_i**: 2
- **stores_t_mu_energy_balance_i**: 4
- **generators_i**: 2867
- **generators_t_p_max_pu_i**: 2639
- **generators_t_p_i**: 2758
- **generators_t_mu_upper_i**: 2763
- **generators_t_mu_lower_i**: 2758
- **generators_t_mu_ramp_limit_up_i**: 2
- **generators_t_mu_ramp_limit_down_i**: 2
- **storage_units_i**: 32
- **storage_units_t_p_i**: 32
- **storage_units_t_p_dispatch_i**: 32
- **storage_units_t_p_store_i**: 32
- **storage_units_t_state_of_charge_i**: 32
- **storage_units_t_mu_upper_i**: 32
- **storage_units_t_mu_lower_i**: 32
- **storage_units_t_mu_energy_balance_i**: 32
- **links_i**: 280
- **links_t_p0_i**: 280
- **links_t_p1_i**: 280
- **links_t_mu_lower_i**: 280
- **links_t_mu_upper_i**: 275
- **carriers_i**: 19
- **buses_i**: 70
- **buses_t_p_i**: 64
- **buses_t_marginal_price_i**: 70
- **global_constraints_i**: 1
- **loads_i**: 58
- **loads_t_p_set_i**: 58
- **loads_t_p_i**: 58

## Global Attributes
- **network__linearized_uc**: 1
- **network__multi_invest**: 1
- **network_name**: 
- **network_objective**: 72727606.62626582
- **network_objective_constant**: 0.0
- **network_pypsa_version**: 0.30.2
- **network_srid**: 4326
- **crs**: {"_crs": "GEOGCRS[\"WGS 84\",ENSEMBLE[\"World Geodetic System 1984 ensemble\",MEMBER[\"World Geodetic System 1984 (Transit)\"],MEMBER[\"World Geodetic System 1984 (G730)\"],MEMBER[\"World Geodetic System 1984 (G873)\"],MEMBER[\"World Geodetic System 1984 (G1150)\"],MEMBER[\"World Geodetic System 1984 (G1674)\"],MEMBER[\"World Geodetic System 1984 (G1762)\"],MEMBER[\"World Geodetic System 1984 (G2139)\"],MEMBER[\"World Geodetic System 1984 (G2296)\"],ELLIPSOID[\"WGS 84\",6378137,298.257223563,LENGTHUNIT[\"metre\",1]],ENSEMBLEACCURACY[2.0]],PRIMEM[\"Greenwich\",0,ANGLEUNIT[\"degree\",0.0174532925199433]],CS[ellipsoidal,2],AXIS[\"geodetic latitude (Lat)\",north,ORDER[1],ANGLEUNIT[\"degree\",0.0174532925199433]],AXIS[\"geodetic longitude (Lon)\",east,ORDER[2],ANGLEUNIT[\"degree\",0.0174532925199433]],USAGE[SCOPE[\"Horizontal component of 3D system.\"],AREA[\"World.\"],BBOX[-90,-180,90,180]],ID[\"EPSG\",4326]]"}
- **meta**: {"run": {"name": "caiso_2025_demo", "disable_progressbar": true, "shared_resources": false, "shared_cutouts": true, "validation": true}, "renewable": {"dataset": "godeeep", "EGS": {"dispatch": "baseload"}, "onwind": {"cutout": "era5", "resource": {"method": "wind", "turbine": "Vestas_V112_3MW", "add_cutout_windspeed": true}, "capacity_per_sqkm": 3, "correction_factor": 1, "corine": {"grid_codes": [20, 30, 40, 60, 100, 111, 112, 113, 114, 115], "distance": 10, "distance_grid_codes": [50]}, "natura": true, "cec": true, "potential": "conservative", "clip_p_max_pu": 0.01, "extendable": true}, "offwind": {"cutout": "era5", "resource": {"method": "wind", "turbine": "NREL_ReferenceTurbine_2020ATB_5.5MW"}, "capacity_per_sqkm": 3, "correction_factor": 1, "corine": {"grid_codes": [80, 200]}, "natura": true, "boem_screen": false, "max_depth": 60, "min_shore_distance": 22000, "max_shore_distance": 65000, "potential": "conservative", "clip_p_max_pu": 0.01, "extendable": true}, "offwind_floating": {"cutout": "era5", "resource": {"method": "wind", "turbine": "NREL_ReferenceTurbine_2020ATB_15MW_offshore", "add_cutout_windspeed": true}, "capacity_per_sqkm": 3, "correction_factor": 1, "corine": {"grid_codes": [80, 200]}, "natura": true, "boem_screen": true, "min_depth": 60, "max_depth": 1300, "min_shore_distance": 22000, "max_shore_distance": 65000, "potential": "conservative", "clip_p_max_pu": 0.01, "extendable": true}, "solar": {"cutout": "era5", "resource": {"method": "pv", "panel": "CSi", "orientation": "latitude_optimal"}, "capacity_per_sqkm": 4.6, "correction_factor": 1, "corine": {"grid_codes": [20, 30, 60, 90, 100]}, "natura": true, "cec": true, "potential": "conservative", "clip_p_max_pu": 0.01, "extendable": true}, "hydro": {"cutout": "era5", "carriers": ["ror", "PHS", "hydro"], "PHS_max_hours": 6, "resource": {"method": "hydro", "hydrobasins": "resources/hybas_na_lev06_v1c.shp", "flowspeed": 1.0}, "hydro_max_hours": "energy_capacity_totals_by_country", "clip_min_inflow": 1.0, "extendable": true, "normalization": {"method": "hydro_capacities", "year": 2013}, "multiplier": 1.1}}, "scenario": {"interconnect": ["western"], "clusters": ["58c"], "simpl": ["300"], "opts": ["REM-3h"], "ll": ["v1.0"], "sector": "E", "planning_horizons": [2025]}, "foresight": "perfect", "model_topology": {"transmission_network": "reeds", "topological_boundaries": "county", "interface_transmission_limits": false, "include": {"reeds_state": ["CA"]}, "aggregate": null}, "enable": {"build_cutout": false}, "renewable_weather_years": [2025], "snapshots": {"start": "2025-07-14 00:00", "end": "2025-07-20 21:00", "inclusive": "both"}, "renewable_scenarios": ["historical"], "renewable_snapshots": {}, "electricity": {"conventional_carriers": ["nuclear", "oil", "OCGT", "CCGT", "coal", "geothermal", "biomass", "waste"], "renewable_carriers": ["onwind", "offwind_floating", "solar", "hydro"], "retirement": "economic", "SAFE_reservemargin": 0.14, "regional_Co2_limits": "config/policy_constraints/regional_Co2_limits.csv", "technology_capacity_targets": "config/policy_constraints/technology_capacity_targets.csv", "portfolio_standards": "config/policy_constraints/portfolio_standards.csv", "SAFE_regional_reservemargins": "config/policy_constraints/SAFE_regional_prm.csv", "transmission_interface_limits": "config/policy_constraints/transmission_interface_limits.csv", "operational_reserve": {"activate": false, "epsilon_load": 0.02, "epsilon_vres": 0.02, "contingency": 4000}, "erm": {"all": 0.15}, "extendable_carriers": {"Generator": [], "StorageUnit": [], "Store": [], "Link": []}, "demand": {"profile": "efs", "scenario": {"efs_case": "reference", "efs_speed": "moderate", "aeo": "reference"}}, "demand_response": {"shift": 0, "marginal_cost": 999999}, "imports": {"enable": true, "costs": 12, "co2_emissions": 0, "capacity_limit": true, "volume_limit": 28, "balancing_period": "year"}, "exports": {"enable": false, "costs": "wholesale", "capacity_limit": true, "volume_limit": "inf", "balancing_period": "month"}, "prm": {"MISO": 0.183, "NPCC_NE": 0.135, "NPCC_NY": 0.15, "PJM": 0.146, "SERC": 0.15, "SPP": 0.16, "ERCOT": 0.1375, "WECC_CA": 0.182, "WECC_NWPP": 0.1435, "WECC_SRSG": 0.1155}}, "conventional": {"unit_commitment": false, "must_run": false, "dynamic_fuel_price": {"enable": false, "pudl": true, "wholesale": true}}, "lines": {"s_max_pu": 0.7, "s_nom_max": Infinity, "max_extension": 0.0, "length_factor": 1.25, "types": {"115.0": "Al/St 240/40 2-bundle 220.0", "138.0": "Al/St 240/40 2-bundle 220.0", "161.0": "Al/St 240/40 2-bundle 220.0", "230.0": "Al/St 240/40 2-bundle 220.0", "345.0": "Al/St 240/40 4-bundle 380.0", "500.0": "Al/St 560/50 4-bundle 750.0", "765.0": "Al/St 560/50 4-bundle 750.0"}}, "links": {"p_max_pu": 1.0, "p_nom_max": Infinity, "max_extension": 0.0}, "co2": {"storage": false, "network": {"enable": false, "capital_cost": 2736000, "marginal_cost": 4, "lifetime": 40, "discount_rate": 0.07}}, "dac": {"enable": false, "granularity": "node", "capital_cost": 6000000, "electricity_input": 2.5, "lifetime": 20, "discount_rate": 0.07}, "costs": {"atb": {"model_case": "Market", "scenario": "Moderate"}, "aeo": {"scenario": "reference"}, "social_discount_rate": 0.02, "ng_fuel_year": 2019, "emission_prices": {"enable": false, "co2": 0.0, "co2_monthly_prices": false}, "ptc_modifier": {"onwind": 27.5, "biomass": 27.5}, "itc_modifier": {"solar": 0.3, "offwind": 0.3, "offwind_floating": 0.3, "EGS": 0.3, "geothermal": 0.3, "SMR": 0.3, "nuclear": 0.3, "hydro": 0.3, "2hr_battery_storage": 0.3, "4hr_battery_storage": 0.3, "6hr_battery_storage": 0.3, "8hr_battery_storage": 0.3, "10hr_battery_storage": 0.3, "8hr_PHS": 0.3, "10hr_PHS": 0.3, "12hr_PHS": 0.3}, "min_year": {"hydrogen_ct": 2040}, "max_growth": null}, "clustering": {"simplify_network": {"weighting_strategy": "demand-capacity", "to_substations": false, "algorithm": "kmeans", "feature": "solar+onwind-time"}, "cluster_network": {"weighting_strategy": "demand-capacity", "algorithm": "kmeans", "feature": "solar+onwind-time", "exclude_carriers": [], "consider_efficiency_classes": false}, "aggregation_strategies": {"generators": {"build_year": "capacity_weighted_average", "lifetime": "capacity_weighted_average", "start_up_cost": "capacity_weighted_average", "min_up_time": "capacity_weighted_average", "min_down_time": "capacity_weighted_average", "ramp_limit_up": "max", "ramp_limit_down": "max", "committable": "any", "vom_cost": "mean", "fuel_cost": "mean", "heat_rate": "mean"}}, "temporal": {"resolution_elec": "3h", "resolution_sector": false}}, "focus_weights": null, "solving": {"options": {"load_shedding": true, "clip_p_max_pu": 0.01, "noisy_costs": true, "seed": 123, "track_iterations": false, "min_iterations": 4, "max_iterations": 6, "transmission_losses": 2, "linearized_unit_commitment": true, "horizon": 8760, "assign_all_duals": true}, "solver": {"name": "mosek", "options": "mosek-default"}, "solver_options": {"highs-default": {"threads": 4, "solver": "ipm", "run_crossover": "on", "small_matrix_value": "1e-6", "large_matrix_value": "1e9", "primal_feasibility_tolerance": "1e-5", "dual_feasibility_tolerance": "1e-5", "ipm_optimality_tolerance": "1e-4", "parallel": "on", "random_seed": 123}, "gurobi-default": {"threads": 8, "method": 2, "crossover": 0, "BarHomogeneous": 1, "BarConvTol": 1e-05, "OptimalityTol": 0.0001, "FeasibilityTol": 0.001, "ScaleFlag": 1, "Seed": 123, "AggFill": 0, "PreDual": 0, "GURO_PAR_BARDENSETHRESH": 200}, "gurobi-numeric-focus": {"name": "gurobi", "NumericFocus": 3, "method": 2, "crossover": 0, "BarHomogeneous": 1, "BarConvTol": 1e-05, "FeasibilityTol": 0.0001, "OptimalityTol": 0.0001, "ObjScale": -0.5, "threads": 8, "Seed": 123}, "gurobi-fallback": {"name": "gurobi", "crossover": 0, "method": 2, "BarHomogeneous": 1, "BarConvTol": 1e-05, "FeasibilityTol": 1e-05, "OptimalityTol": 1e-05, "Seed": 123, "threads": 8}, "cplex-default": {"threads": 4, "lpmethod": 4, "solutiontype": 2, "barrier.convergetol": 1e-05, "feasopt.tolerance": 1e-06}, "cbc-default": {}, "glpk-default": {}, "mosek-default": {}}, "mem": 30000, "walltime": "12:00:00"}, "walltime": {"build_renewable_profiles": "04:00:00", "build_fuel_prices": "00:20:00", "add_demand": "02:00:00", "add_electricity": "04:00:00", "simplify_network": "05:00:00", "cluster_network": "04:00:00", "solve_network": "20:00:00"}, "custom_files": {"activate": false, "files_path": "", "network_name": ""}, "pudl_path": "s3://pudl.catalyst.coop/v2026.5.0", "pudl_fleet_max_year": 2024, "api": {"eia": "8DlG3YK2kP4FCgZMGiNhf6cRRCUr1p5vrkHdwyei"}, "__default__": {"account": null, "partition": null, "email": null, "walltime": "00:30:00", "cpus_per_task": 1, "chdir": null, "output": "logs/{rule}/log-%j.out", "error": "logs/{rule}/errlog-%j.err"}, "build_renewable_profiles": {"walltime": "06:00:00"}, "add_electricity": {"walltime": "06:00:00"}, "simplify_network": {"walltime": "09:00:00"}, "cluster_network": {"walltime": "09:00:00"}, "solve_network": {"walltime": "24:00:00"}, "solve_network_validation": {"walltime": "09:00:00"}, "add_demand": {"walltime": "02:00:00"}, "atlite": {"default_cutout": "era5_2019", "nprocesses": 8, "show_progress": false, "cutouts": {"era5_2019": {"module": "era5", "time": ["2019", "2019"]}}, "interconnects": {"western": {"x": [-126, -99], "y": [27, 50], "dx": 0.3, "dy": 0.3}, "eastern": {"x": [-109, -65], "y": [23, 50], "dx": 0.3, "dy": 0.3}, "texas": {"x": [-110, -90], "y": [24, 37], "dx": 0.3, "dy": 0.3}, "usa": {"x": [-126, -65], "y": [23, 50], "dx": 0.3, "dy": 0.3}}}, "offshore_shape": {"use": "eez"}, "offshore_network": {"bus_spacing": 25000}, "plotting": {"costs_max": 800, "costs_threshold": 1, "energy_max": 15000.0, "energy_min": -10000.0, "energy_threshold": 50.0, "tech_colors": {"onwind": "#235ebc", "wind": "#235ebc", "onshore wind": "#235ebc", "offwind": "#dd6895", "offshore wind": "#6895dd", "offwind-ac": "#6895dd", "offshore wind ac": "#6895dd", "offwind-dc": "#74c6f2", "offshore wind dc": "#74c6f2", "offwind_floating": "#11a1c1", "hydro": "#08ad97", "hydro+PHS": "#08ad97", "PHS": "#08ad97", "hydro reservoir": "#08ad97", "hydroelectricity": "#08ad97", "ror": "#4adbc8", "run of river": "#4adbc8", "solar": "#f9d002", "solar PV": "#f9d002", "solar thermal": "#ffef60", "biomass": "#0c6013", "solid biomass": "#06540d", "biogas": "#23932d", "waste": "#68896b", "geothermal": "#ba91b1", "OCGT": "#d35050", "gas": "#d35050", "ng": "#d35050", "natural gas": "#d35050", "CCGT": "#b20101", "nuclear": "#ff9000", "coal": "#707070", "lignite": "#9e5a01", "oil": "#262626", "H2": "#ea048a", "hydrogen storage": "#ea048a", "battery": "#b8ea04", "2hr_battery_storage": "#aee000", "4hr_battery_storage": "#a4d600", "6hr_battery_storage": "#9acc00", "8hr_battery_storage": "#90c200", "10hr_battery_storage": "#86b800", "Electric load": "#f9d002", "electricity": "#f9d002", "lines": "#70af1d", "transmission lines": "#70af1d", "AC-AC": "#70af1d", "AC line": "#70af1d", "AC": "#70af1d", "links": "#8a1caf", "HVDC links": "#8a1caf", "DC-DC": "#8a1caf", "DC link": "#8a1caf", "DC": "#8a1caf", "Load": "#2ad55f", "imports": "#9f2ad5", "exports": "#B79B48", "res-elec": "#f9d002", "res-heat": "#E79CA2", "res-cool": "#9CE7E2", "com-elec": "#f9d002", "com-heat": "#E79CA2", "com-cool": "#9CE7E2", "ind-elec": "#f9d002", "ind-heat": "#E79CA2", "trn-elec": "#f9d002", "coal-95CCS": "#4b4b4b", "coal-99CCS": "#2e2e2e", "coal-95CC": "#4b4b4b", "coal-99CC": "#2e2e2e", "SMR": "#ff5733", "CCGT-95CCS": "#800000", "CCGT-95CC": "#800000", "8hr_PHS": "#069686", "10hr_PHS": "#058a79", "12hr_PHS": "#047d6c", "8hr_PHS_discharger": "#069686", "10hr_PHS_discharger": "#058a79", "12hr_PHS_discharger": "#047d6c", "8hr_PHS_charger": "#069686", "10hr_PHS_charger": "#058a79", "12hr_PHS_charger": "#047d6c", "hydrogen_ct": "#ea048a", "demand_response": "#8c03fc", "dac": "#4c004c", "lpg": "#70217b", "ch4": "#539307", "co2": "#cd7e0d", "res-total-elec": "#f9d002", "res-urban-elec": "#f9d002", "res-rural-elec": "#f9d002", "res-total-heat": "#E79CA2", "res-urban-heat": "#E79CA2", "res-rural-heat": "#E79CA2", "res-total-cool": "#196CE6", "res-urban-cool": "#196CE6", "res-rural-cool": "#196CE6", "com-total-elec": "#f9d002", "com-urban-elec": "#f9d002", "com-rural-elec": "#f9d002", "com-total-heat": "#E79CA2", "com-urban-heat": "#E79CA2", "com-rural-heat": "#E79CA2", "com-total-cool": "#196CE6", "com-urban-cool": "#196CE6", "com-rural-cool": "#196CE6", "gas storage": "#f69d09", "gas pipeline": "#f69d09", "gas trade": "ae3dc2", "gas production": "#d35050", "trn-veh": "#0a0100", "trn-veh-lgt": "#0a0100", "trn-veh-med": "#0a0100", "trn-veh-hvy": "#0a0100", "trn-veh-bus": "#0a0100", "trn-elec-veh": "#0a0100", "trn-elec-veh-lgt": "#2BAAD4", "trn-elec-veh-med": "#2FD085", "trn-elec-veh-hvy": "#6DCD32", "trn-elec-veh-bus": "#d3d32c", "trn-lpg": "#0a0100", "trn-lpg-veh": "#0a0100", "trn-lpg-veh-lgt": "#D4552B", "trn-lpg-veh-med": "#D02F7A", "trn-lpg-veh-hvy": "#9232CD", "trn-lpg-veh-bus": "#2C2CD3", "trn-rail": "#0a0100", "trn-rail-psg": "#0a0100", "trn-rail-ship": "#0a0100", "trn-lpg-rail": "#0a0100", "trn-lpg-rail-psg": "#9F9160", "trn-lpg-rail-ship": "#606E9F", "trn-air": "#0a0100", "trn-air-psg": "#0a0100", "trn-lpg-air": "#0a0100", "trn-lpg-air-psg": "#A45B75", "trn-boat": "#0a0100", "trn-boat-ship": "#0a0100", "trn-lpg-boat": "#0a0100", "trn-lpg-boat-ship": "#5BA48A", "res-space-heat": "#F10E1B", "com-space-heat": "#F10E1B", "res-water-heat": "#F10E1B", "com-water-heat": "#F10E1B", "res-rural-space-heat": "#F10E1B", "com-rural-space-heat": "#F10E1B", "res-urban-space-heat": "#F10E1B", "com-urban-space-heat": "#F10E1B", "res-total-space-heat": "#F10E1B", "com-total-space-heat": "#F10E1B", "res-rural-water-heat": "#E817C3", "com-rural-water-heat": "#E817C3", "res-urban-water-heat": "#E817C3", "com-urban-water-heat": "#E817C3", "res-total-water-heat": "#E817C3", "com-total-water-heat": "#E817C3", "res-rural-air-con": "#3B98C4", "com-rural-air-con": "#3B98C4", "res-urban-air-con": "#3B98C4", "com-urban-air-con": "#3B98C4", "res-total-air-con": "#3B98C4", "com-total-air-con": "#3B98C4", "res-rural-gas-furnace": "#F68D09", "com-rural-gas-furnace": "#F68D09", "res-urban-gas-furnace": "#F68D09", "com-urban-gas-furnace": "#F68D09", "res-total-gas-furnace": "#F68D09", "com-total-gas-furnace": "#F68D09", "res-rural-space-gas-furnace": "#F68D09", "com-rural-space-gas-furnace": "#F68D09", "res-urban-space-gas-furnace": "#F68D09", "com-urban-space-gas-furnace": "#F68D09", "res-total-space-gas-furnace": "#F68D09", "com-total-space-gas-furnace": "#F68D09", "res-rural-oil-furnace": "#006B88", "com-rural-oil-furnace": "#006B88", "res-urban-oil-furnace": "#006B88", "com-urban-oil-furnace": "#006B88", "res-total-oil-furnace": "#006B88", "com-total-oil-furnace": "#006B88", "res-rural-space-oil-furnace": "#006B88", "com-rural-space-oil-furnace": "#006B88", "com-urban-space-oil-furnace": "#006B88", "res-urban-space-oil-furnace": "#006B88", "com-total-space-oil-furnace": "#006B88", "res-total-space-oil-furnace": "#006B88", "res-rural-elec-furnace": "#DCEC13", "com-rural-elec-furnace": "#DCEC13", "res-urban-elec-furnace": "#DCEC13", "com-urban-elec-furnace": "#DCEC13", "res-total-elec-furnace": "#DCEC13", "com-total-elec-furnace": "#DCEC13", "res-rural-space-elec-furnace": "#DCEC13", "com-rural-space-elec-furnace": "#DCEC13", "com-urban-space-elec-furnace": "#DCEC13", "res-urban-space-elec-furnace": "#DCEC13", "com-total-space-elec-furnace": "#DCEC13", "res-total-space-elec-furnace": "#DCEC13", "res-rural-water-gas": "#D87627", "com-rural-water-gas": "#D87627", "res-urban-water-gas": "#D87627", "com-urban-water-gas": "#D87627", "res-total-water-gas": "#D87627", "com-total-water-gas": "#D87627", "res-rural-water-oil": "#774115", "com-rural-water-oil": "#774115", "res-urban-water-oil": "#774115", "com-urban-water-oil": "#774115", "res-total-water-oil": "#774115", "com-total-water-oil": "#774115", "res-rural-water-elec": "#5CDC23", "com-rural-water-elec": "#5CDC23", "res-urban-water-elec": "#5CDC23", "com-urban-water-elec": "#5CDC23", "res-total-water-elec": "#5CDC23", "com-total-water-elec": "#5CDC23", "res-rural-ashp": "#E14C1E", "com-rural-ashp": "#E14C1E", "res-urban-ashp": "#E14C1E", "com-urban-ashp": "#E14C1E", "res-total-ashp": "#E14C1E", "com-total-ashp": "#E14C1E", "res-rural-gshp": "#D926A6", "com-rural-gshp": "#D926A6", "res-urban-gshp": "#D926A6", "com-urban-gshp": "#D926A6", "res-total-gshp": "#D926A6", "com-total-gshp": "#D926A6", "res-rural-space-heat-store": "#3096CF", "com-rural-space-heat-store": "#3096CF", "res-urban-space-heat-store": "#3096CF", "com-urban-space-heat-store": "#3096CF", "res-total-space-heat-store": "#3096CF", "com-total-space-heat-store": "#3096CF", "res-elec-infra": "#9CD926", "com-elec-infra": "#9CD926", "ind-coal-furnace": "#EC13B4", "ind-heat-pump": "#E83E17", "ind-gas-furnace": "#E0AE1F", "ind-elec-infra": "#9CD926", "res-total-elec-dr": "#45ba75", "res-urban-elec-dr": "#45ba75", "res-rural-elec-dr": "#45ba75", "res-total-heat-dr": "#BA458A", "res-urban-heat-dr": "#BA458A", "res-rural-heat-dr": "#BA458A", "res-total-space-heat-dr": "#A55A80", "res-urban-space-heat-dr": "#A55A80", "res-rural-space-heat-dr": "#A55A80", "res-total-water-heat-dr": "#A844BB", "res-urban-water-heat-dr": "#A844BB", "res-rural-water-heat-dr": "#A844BB", "res-total-cool-dr": "#4692B9", "res-urban-cool-dr": "#4692B9", "res-rural-cool-dr": "#4692B9", "com-total-elec-dr": "#45ba75", "com-urban-elec-dr": "#45ba75", "com-rural-elec-dr": "#45ba75", "com-total-heat-dr": "#BA458A", "com-urban-heat-dr": "#BA458A", "com-rural-heat-dr": "#BA458A", "com-total-space-heat-dr": "#A55A80", "com-urban-space-heat-dr": "#A55A80", "com-rural-space-heat-dr": "#A55A80", "com-total-water-heat-dr": "#A844BB", "com-urban-water-heat-dr": "#A844BB", "com-rural-water-heat-dr": "#A844BB", "com-total-cool-dr": "#4692B9", "com-urban-cool-dr": "#4692B9", "com-rural-cool-dr": "#4692B9", "ind-elec-dr": "#45ba75", "ind-heat-dr": "#BA458A", "trn-elec-dr": "#45ba75", "trn-elec-veh-dr": "#45ba75"}, "nice_names": {"OCGT": "Open-Cycle Gas", "CCGT": "Combined-Cycle Gas", "offwind": "Fixed Bottom Offshore Wind", "offwind_floating": "Floating Offshore Wind", "onwind": "Onshore Wind", "solar": "Solar", "PHS": "Pumped Hydro Storage", "hydro": "Reservoir & Dam", "battery": "Battery Storage", "H2": "Hydrogen Storage", "lines": "Transmission Lines", "ror": "Run of River", "Load": "Load Shed", "hydrogen_ct": "Hydrogen Combustion Turbine", "demand_response": "Demand Response", "dac": "Direct Air Capture", "imports": "Electricity Imports", "exports": "Electricity Exports", "lpg": "Motor Oil", "ch4": "Methane", "co2": "Carbon Dioxide", "res-elec": "Residential Electrical", "res-total-elec": "Residential Electrical", "res-urban-elec": "Residential Urban Electrical", "res-rural-elec": "Residential Rural Electrical", "res-heat": "Residential Heating", "res-total-heat": "Residential Heating", "res-urban-heat": "Residential Urban Heating", "res-rural-heat": "Residential Rural Heating", "res-cool": "Residential Cooling", "res-total-cool": "Cool", "res-urban-cool": "Cool", "res-rural-cool": "Cool", "com-elec": "Commercial Electrical", "com-total-elec": "Commercial Electrical", "com-urban-elec": "Commercial Electrical", "com-rural-elec": "Commercial Electrical", "com-heat": "Commercial Heating", "com-total-heat": "Commercial Heating", "com-urban-heat": "Commercial Urban Heating", "com-rural-heat": "Commercial Rural Heating", "com-cool": "Commercial Cooling", "com-total-cool": "Cool", "com-urban-cool": "Cool", "com-rural-cool": "Cool", "ind-elec": "Industrial Electrical", "ind-heat": "Industrial Heating", "trn-elec": "Transportation Electrical", "gas storage": "Natural Gas Storage", "gas pipeline": "Natural Gas Pipeline", "gas trade": "Natural Gas Trading", "gas production": "Natural Gas Production", "trn-veh": "Vehicles", "trn-veh-lgt": "Light-Duty Vehicles", "trn-veh-med": "Medium-Duty Vehicles", "trn-veh-hvy": "Heavy-Duty Vehicles", "trn-veh-bus": "Buses", "trn-elec-veh": "Electric Vehicles", "trn-elec-veh-lgt": "Light Duty EV", "trn-elec-veh-med": "Medium Duty EV", "trn-elec-veh-hvy": "Heavy Duty EV", "trn-elec-veh-bus": "Electric Bus", "trn-lpg": "Transportaion Gas", "trn-lpg-veh": "Gas Vehicles", "trn-lpg-veh-lgt": "Light Duty ICE", "trn-lpg-veh-med": "Medium Duty ICE", "trn-lpg-veh-hvy": "Heavy Duty ICE", "trn-lpg-veh-bus": "Gas Bus", "trn-rail": "Rail", "trn-rail-psg": "Passenger Rail", "trn-rail-ship": "Shipping Rail", "trn-lpg-rail": "Rail Oil", "trn-lpg-rail-psg": "Passenger Rail Oil", "trn-lpg-rail-ship": "Shipping Rail Oil", "trn-air": "Airplane", "trn-air-psg": "Passenger Airplane", "trn-lpg-air": "Airplane Gas", "trn-lpg-air-psg": "Passenger Air Gas", "trn-boat": "Marine Shipping", "trn-boat-ship": "Marine Shipping", "trn-lpg-boat": "Domestic Marine Shipping Gas", "trn-lpg-boat-ship": "Domestic Marine Shipping Gas", "res-space-heat": "Space Heat", "com-space-heat": "Space Heat", "res-water-heat": "Water Heat", "com-water-heat": "Water Heat", "res-rural-space-heat": "Space Heat", "com-rural-space-heat": "Space Heat", "res-urban-space-heat": "Space Heat", "com-urban-space-heat": "Space Heat", "res-total-space-heat": "Space Heat", "com-total-space-heat": "Space Heat", "res-rural-water-heat": "Water Heat", "com-rural-water-heat": "Water Heat", "res-urban-water-heat": "Water Heat", "com-urban-water-heat": "Water Heat", "res-total-water-heat": "Water Heat", "com-total-water-heat": "Water Heat", "res-rural-air-con": "Air Conditioner", "com-rural-air-con": "Air Conditioner", "res-urban-air-con": "Air Conditioner", "com-urban-air-con": "Air Conditioner", "res-total-air-con": "Air Conditioner", "com-total-air-con": "Air Conditioner", "res-rural-gas-furnace": "Gas Furnace", "com-rural-gas-furnace": "Gas Furnace", "res-urban-gas-furnace": "Gas Furnace", "com-urban-gas-furnace": "Gas Furnace", "res-total-gas-furnace": "Gas Furnace", "com-total-gas-furnace": "Gas Furnace", "res-rural-space-gas-furnace": "Gas Furnace", "com-rural-space-gas-furnace": "Gas Furnace", "res-urban-space-gas-furnace": "Gas Furnace", "com-urban-space-gas-furnace": "Gas Furnace", "res-total-space-gas-furnace": "Gas Furnace", "com-total-space-gas-furnace": "Gas Furnace", "res-rural-oil-furnace": "Oil Furnace", "com-rural-oil-furnace": "Oil Furnace", "res-urban-oil-furnace": "Oil Furnace", "com-urban-oil-furnace": "Oil Furnace", "res-total-oil-furnace": "Oil Furnace", "com-total-oil-furnace": "Oil Furnace", "res-rural-space-oil-furnace": "Oil Furnace", "com-rural-space-oil-furnace": "Oil Furnace", "com-urban-space-oil-furnace": "Oil Furnace", "res-urban-space-oil-furnace": "Oil Furnace", "com-total-space-oil-furnace": "Oil Furnace", "res-total-space-oil-furnace": "Oil Furnace", "res-rural-elec-furnace": "Electric Furnace", "com-rural-elec-furnace": "Electric Furnace", "res-urban-elec-furnace": "Electric Furnace", "com-urban-elec-furnace": "Electric Furnace", "res-total-elec-furnace": "Electric Furnace", "com-total-elec-furnace": "Electric Furnace", "res-rural-space-elec-furnace": "Electric Furnace", "com-rural-space-elec-furnace": "Electric Furnace", "com-urban-space-elec-furnace": "Electric Furnace", "res-urban-space-elec-furnace": "Electric Furnace", "com-total-space-elec-furnace": "Electric Furnace", "res-total-space-elec-furnace": "Electric Furnace", "res-rural-water-gas": "Gas Water Heater", "com-rural-water-gas": "Gas Water Heater", "res-urban-water-gas": "Gas Water Heater", "com-urban-water-gas": "Gas Water Heater", "res-total-water-gas": "Gas Water Heater", "com-total-water-gas": "Gas Water Heater", "res-rural-water-oil": "Oil Water Heater", "com-rural-water-oil": "Oil Water Heater", "res-urban-water-oil": "Oil Water Heater", "com-urban-water-oil": "Oil Water Heater", "res-total-water-oil": "Oil Water Heater", "com-total-water-oil": "Oil Water Heater", "res-rural-water-elec": "Electric Water Heater", "com-rural-water-elec": "Electric Water Heater", "res-urban-water-elec": "Electric Water Heater", "com-urban-water-elec": "Electric Water Heater", "res-total-water-elec": "Electric Water Heater", "com-total-water-elec": "Electric Water Heater", "res-rural-ashp": "Air Source Heat Pump", "com-rural-ashp": "Air Source Heat Pump", "res-urban-ashp": "Air Source Heat Pump", "com-urban-ashp": "Air Source Heat Pump", "res-total-ashp": "Air Source Heat Pump", "com-total-ashp": "Air Source Heat Pump", "res-rural-gshp": "Ground Source Heat Pump", "com-rural-gshp": "Ground Source Heat Pump", "res-urban-gshp": "Ground Source Heat Pump", "com-urban-gshp": "Ground Source Heat Pump", "res-total-gshp": "Ground Source Heat Pump", "com-total-gshp": "Ground Source Heat Pump", "res-rural-space-heat-store": "Building Insulation", "com-rural-space-heat-store": "Building Insulation", "res-urban-space-heat-store": "Building Insulation", "com-urban-space-heat-store": "Building Insulation", "res-total-space-heat-store": "Building Insulation", "com-total-space-heat-store": "Building Insulation", "res-elec-infra": "Electric Distribution", "com-elec-infra": "Electric Distribution", "ind-coal-furnace": "Coal furnace", "ind-heat-pump": "Heat Pump", "ind-gas-furnace": "Gas Furnace", "ind-elec-infra": "Electric Distribution", "res-total-elec-dr": "Electric Demand Response", "res-urban-elec-dr": "Urban Electric Demand Response", "res-rural-elec-dr": "Rural Electric Demand Response", "res-total-heat-dr": "Heating Demand Response", "res-urban-heat-dr": "Urban Heating Demand Response", "res-rural-heat-dr": "Rural Heating Demand Response", "res-total-space-heat-dr": "Space Heating Demand Response", "res-urban-space-heat-dr": "Urban Space Heating Demand Response", "res-rural-space-heat-dr": "Rural Space Heating Demand Response", "res-total-water-heat-dr": "Water Heating Demand Response", "res-urban-water-heat-dr": "Urban Water Heating Demand Response", "res-rural-water-heat-dr": "Rural Water Heating Demand Response", "res-total-cool-dr": "Cooling Demand Response", "res-urban-cool-dr": "Urban Cooling Demand Response", "res-rural-cool-dr": "Rural Cooling Demand Response", "com-total-elec-dr": "Electric Demand Response", "com-urban-elec-dr": "Urban Electric Demand Response", "com-rural-elec-dr": "Rural Electric Demand Response", "com-total-heat-dr": "Heating Demand Response", "com-urban-heat-dr": "Urban Heating Demand Response", "com-rural-heat-dr": "Rural Heating Demand Response", "com-total-space-heat-dr": "Space Heating Demand Response", "com-urban-space-heat-dr": "Urban Space Heating Demand Response", "com-rural-space-heat-dr": "Rural Space Heating Demand Response", "com-total-water-heat-dr": "Water Heating Demand Response", "com-urban-water-heat-dr": "Urban Water Heating Demand Response", "com-rural-water-heat-dr": "Rural Water Heating Demand Response", "com-total-cool-dr": "Cooling Demand Response", "com-urban-cool-dr": "Urban Cooling Demand Response", "com-rural-cool-dr": "Rural Cooling Demand Response", "ind-elec-dr": "Electric Demand Response", "ind-heat-dr": "Heating Demand Response", "trn-elec-dr": "Electric Demand Response", "trn-elec-veh-dr": "Electric Demand Response"}}, "sector": {"co2": {"sequestration_potential": 0, "policy": "config/policy_constraints/sector_co2_limits.csv"}, "natural_gas": {"imports": {"min": 0.99, "max": 1.01}, "exports": {"min": 0.99, "max": 1.01}, "cyclic_storage": true, "standing_loss": 0, "marginal_cost_multiplier": 1, "existing_pipeline_multiplier": 1}, "methane": {"upstream_leakage_rate": 0.02, "downstream_leakage_rate": 0.02, "gwp": 0}, "heating": {"heat_pump_sink_T": 55.0}, "service_sector": {"dynamic_costs": true, "split_res_com": true, "split_urban_rural": false, "water_heating": {"split_space_water": true, "simple_storage": true, "n_hours": 4}, "split_space_water_heating": true, "brownfield": true, "scale_exising_stock": true, "gas_connection": {"rural": 1, "urban": 1}, "technologies": {"space_heating": {"elec_furnace": true, "gas_furnace": true, "oil_furnace": true, "heat_pump": true, "air_con": true}, "water_heating": {"elec_water_tank": true, "gas_water_tank": true, "oil_water_tank": false}, "standing_losses": {"space": 0.05, "water": 0.01}}, "loads": {"heating": true, "cooling": true}, "demand_response": {"shift": 0, "marginal_cost": {"electricity": 5, "space-heat": 5, "cool": 5, "heat": 5}}}, "transport_sector": {"brownfield": true, "dynamic_costs": true, "ev_policy": "config/policy_constraints/ev_policy.csv", "must_run_evs": true, "modes": {"vehicle": true, "rail": true, "air": true, "boat": true}, "demand_response": {"shift": 0, "marginal_cost": 10}}, "industrial_sector": {"brownfield": true, "dynamic_costs": true, "technologies": {"gas_furnace": true, "coal_furnace": true, "heat_pump": true}, "min_fossil_generation": 66, "demand_response": {"shift": 0, "marginal_cost": {"electricity": 5, "heat": 5}}}}, "wildcards": {"interconnect": "western", "simpl": "300", "clusters": "58c", "ll": "v1.0", "opts": "REM-3h", "sector": "E"}}

## Variables
### snapshots
- **Type**: int64
- **Shape**: (56,)
- **Dimensions**: snapshots(56)

### snapshots_period
- **Type**: int32
- **Shape**: (56,)
- **Dimensions**: snapshots(56)

### snapshots_timestep
- **Type**: int64
- **Shape**: (56,)
- **Dimensions**: snapshots(56)
- **Attributes**:
  - units: hours since 2025-07-14 00:00:00
  - calendar: proleptic_gregorian

### snapshots_objective
- **Type**: float64
- **Shape**: (56,)
- **Dimensions**: snapshots(56)
- **Attributes**:
  - _FillValue: nan

### snapshots_generators
- **Type**: float64
- **Shape**: (56,)
- **Dimensions**: snapshots(56)
- **Attributes**:
  - _FillValue: nan

### snapshots_stores
- **Type**: float64
- **Shape**: (56,)
- **Dimensions**: snapshots(56)
- **Attributes**:
  - _FillValue: nan

### investment_periods
- **Type**: int64
- **Shape**: (1,)
- **Dimensions**: investment_periods(1)

### investment_periods_objective
- **Type**: float64
- **Shape**: (1,)
- **Dimensions**: investment_periods(1)
- **Attributes**:
  - _FillValue: nan

### investment_periods_years
- **Type**: float64
- **Shape**: (1,)
- **Dimensions**: investment_periods(1)
- **Attributes**:
  - _FillValue: nan

### stores_i
- **Type**: <class 'str'>
- **Shape**: (12,)
- **Dimensions**: stores_i(12)

### stores_bus
- **Type**: <class 'str'>
- **Shape**: (12,)
- **Dimensions**: stores_i(12)

### stores_carrier
- **Type**: <class 'str'>
- **Shape**: (12,)
- **Dimensions**: stores_i(12)

### stores_e_nom_extendable
- **Type**: int8
- **Shape**: (12,)
- **Dimensions**: stores_i(12)
- **Attributes**:
  - dtype: bool

### stores_e_nom_max
- **Type**: float64
- **Shape**: (12,)
- **Dimensions**: stores_i(12)
- **Attributes**:
  - _FillValue: nan

### stores_e_min_pu
- **Type**: float64
- **Shape**: (12,)
- **Dimensions**: stores_i(12)
- **Attributes**:
  - _FillValue: nan

### stores_e_max_pu
- **Type**: float64
- **Shape**: (12,)
- **Dimensions**: stores_i(12)
- **Attributes**:
  - _FillValue: nan

### stores_e_cyclic_per_period
- **Type**: int8
- **Shape**: (12,)
- **Dimensions**: stores_i(12)
- **Attributes**:
  - dtype: bool

### stores_marginal_cost
- **Type**: float64
- **Shape**: (12,)
- **Dimensions**: stores_i(12)
- **Attributes**:
  - _FillValue: nan

### stores_e_nom_opt
- **Type**: float64
- **Shape**: (12,)
- **Dimensions**: stores_i(12)
- **Attributes**:
  - _FillValue: nan

### stores_t_p_i
- **Type**: <class 'str'>
- **Shape**: (12,)
- **Dimensions**: stores_t_p_i(12)

### stores_t_p
- **Type**: float64
- **Shape**: (56, 12)
- **Dimensions**: snapshots(56), stores_t_p_i(12)
- **Attributes**:
  - _FillValue: nan

### stores_t_e_i
- **Type**: <class 'str'>
- **Shape**: (12,)
- **Dimensions**: stores_t_e_i(12)

### stores_t_e
- **Type**: float64
- **Shape**: (56, 12)
- **Dimensions**: snapshots(56), stores_t_e_i(12)
- **Attributes**:
  - _FillValue: nan

### stores_t_mu_lower_i
- **Type**: <class 'str'>
- **Shape**: (2,)
- **Dimensions**: stores_t_mu_lower_i(2)

### stores_t_mu_lower
- **Type**: float64
- **Shape**: (56, 2)
- **Dimensions**: snapshots(56), stores_t_mu_lower_i(2)
- **Attributes**:
  - _FillValue: nan

### stores_t_mu_energy_balance_i
- **Type**: <class 'str'>
- **Shape**: (4,)
- **Dimensions**: stores_t_mu_energy_balance_i(4)

### stores_t_mu_energy_balance
- **Type**: float64
- **Shape**: (56, 4)
- **Dimensions**: snapshots(56), stores_t_mu_energy_balance_i(4)
- **Attributes**:
  - _FillValue: nan

### generators_i
- **Type**: <class 'str'>
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)

### generators_carrier
- **Type**: <class 'str'>
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)

### generators_bus
- **Type**: <class 'str'>
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)

### generators_p_nom_min
- **Type**: float64
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)
- **Attributes**:
  - _FillValue: nan

### generators_p_nom
- **Type**: float64
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)
- **Attributes**:
  - _FillValue: nan

### generators_ramp_limit_up
- **Type**: float64
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)
- **Attributes**:
  - _FillValue: nan

### generators_ramp_limit_down
- **Type**: float64
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)
- **Attributes**:
  - _FillValue: nan

### generators_efficiency
- **Type**: float64
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)
- **Attributes**:
  - _FillValue: nan

### generators_marginal_cost
- **Type**: float64
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)
- **Attributes**:
  - _FillValue: nan

### generators_capital_cost
- **Type**: float64
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)
- **Attributes**:
  - _FillValue: nan

### generators_build_year
- **Type**: int64
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)

### generators_lifetime
- **Type**: float64
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)
- **Attributes**:
  - _FillValue: nan

### generators_control
- **Type**: <class 'str'>
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)

### generators_p_nom_max
- **Type**: float64
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)
- **Attributes**:
  - _FillValue: nan

### generators_p_max_pu
- **Type**: float64
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)
- **Attributes**:
  - _FillValue: nan

### generators_weight
- **Type**: float64
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)
- **Attributes**:
  - _FillValue: nan

### generators_vom_cost
- **Type**: float64
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)
- **Attributes**:
  - _FillValue: nan

### generators_fuel_cost
- **Type**: float64
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)
- **Attributes**:
  - _FillValue: nan

### generators_heat_rate
- **Type**: float64
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)
- **Attributes**:
  - _FillValue: nan

### generators_carrier_base
- **Type**: <class 'str'>
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)

### generators_land_region
- **Type**: <class 'str'>
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)

### generators_sign
- **Type**: float64
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)
- **Attributes**:
  - _FillValue: nan

### generators_p_nom_opt
- **Type**: float64
- **Shape**: (2867,)
- **Dimensions**: generators_i(2867)
- **Attributes**:
  - _FillValue: nan

### generators_t_p_max_pu_i
- **Type**: <class 'str'>
- **Shape**: (2639,)
- **Dimensions**: generators_t_p_max_pu_i(2639)

### generators_t_p_max_pu
- **Type**: float64
- **Shape**: (56, 2639)
- **Dimensions**: snapshots(56), generators_t_p_max_pu_i(2639)
- **Attributes**:
  - _FillValue: nan

### generators_t_p_i
- **Type**: <class 'str'>
- **Shape**: (2758,)
- **Dimensions**: generators_t_p_i(2758)

### generators_t_p
- **Type**: float64
- **Shape**: (56, 2758)
- **Dimensions**: snapshots(56), generators_t_p_i(2758)
- **Attributes**:
  - _FillValue: nan

### generators_t_mu_upper_i
- **Type**: <class 'str'>
- **Shape**: (2763,)
- **Dimensions**: generators_t_mu_upper_i(2763)

### generators_t_mu_upper
- **Type**: float64
- **Shape**: (56, 2763)
- **Dimensions**: snapshots(56), generators_t_mu_upper_i(2763)
- **Attributes**:
  - _FillValue: nan

### generators_t_mu_lower_i
- **Type**: <class 'str'>
- **Shape**: (2758,)
- **Dimensions**: generators_t_mu_lower_i(2758)

### generators_t_mu_lower
- **Type**: float64
- **Shape**: (56, 2758)
- **Dimensions**: snapshots(56), generators_t_mu_lower_i(2758)
- **Attributes**:
  - _FillValue: nan

### generators_t_mu_ramp_limit_up_i
- **Type**: <class 'str'>
- **Shape**: (2,)
- **Dimensions**: generators_t_mu_ramp_limit_up_i(2)

### generators_t_mu_ramp_limit_up
- **Type**: float64
- **Shape**: (56, 2)
- **Dimensions**: snapshots(56), generators_t_mu_ramp_limit_up_i(2)
- **Attributes**:
  - _FillValue: nan

### generators_t_mu_ramp_limit_down_i
- **Type**: <class 'str'>
- **Shape**: (2,)
- **Dimensions**: generators_t_mu_ramp_limit_down_i(2)

### generators_t_mu_ramp_limit_down
- **Type**: float64
- **Shape**: (56, 2)
- **Dimensions**: snapshots(56), generators_t_mu_ramp_limit_down_i(2)
- **Attributes**:
  - _FillValue: nan

### storage_units_i
- **Type**: <class 'str'>
- **Shape**: (32,)
- **Dimensions**: storage_units_i(32)

### storage_units_carrier
- **Type**: <class 'str'>
- **Shape**: (32,)
- **Dimensions**: storage_units_i(32)

### storage_units_bus
- **Type**: <class 'str'>
- **Shape**: (32,)
- **Dimensions**: storage_units_i(32)

### storage_units_p_nom
- **Type**: float64
- **Shape**: (32,)
- **Dimensions**: storage_units_i(32)
- **Attributes**:
  - _FillValue: nan

### storage_units_p_nom_max
- **Type**: float64
- **Shape**: (32,)
- **Dimensions**: storage_units_i(32)
- **Attributes**:
  - _FillValue: nan

### storage_units_capital_cost
- **Type**: float64
- **Shape**: (32,)
- **Dimensions**: storage_units_i(32)
- **Attributes**:
  - _FillValue: nan

### storage_units_max_hours
- **Type**: float64
- **Shape**: (32,)
- **Dimensions**: storage_units_i(32)
- **Attributes**:
  - _FillValue: nan

### storage_units_efficiency_store
- **Type**: float64
- **Shape**: (32,)
- **Dimensions**: storage_units_i(32)
- **Attributes**:
  - _FillValue: nan

### storage_units_efficiency_dispatch
- **Type**: float64
- **Shape**: (32,)
- **Dimensions**: storage_units_i(32)
- **Attributes**:
  - _FillValue: nan

### storage_units_cyclic_state_of_charge
- **Type**: int8
- **Shape**: (32,)
- **Dimensions**: storage_units_i(32)
- **Attributes**:
  - dtype: bool

### storage_units_control
- **Type**: <class 'str'>
- **Shape**: (32,)
- **Dimensions**: storage_units_i(32)

### storage_units_p_min_pu
- **Type**: float64
- **Shape**: (32,)
- **Dimensions**: storage_units_i(32)
- **Attributes**:
  - _FillValue: nan

### storage_units_p_max_pu
- **Type**: float64
- **Shape**: (32,)
- **Dimensions**: storage_units_i(32)
- **Attributes**:
  - _FillValue: nan

### storage_units_marginal_cost
- **Type**: float64
- **Shape**: (32,)
- **Dimensions**: storage_units_i(32)
- **Attributes**:
  - _FillValue: nan

### storage_units_p_nom_opt
- **Type**: float64
- **Shape**: (32,)
- **Dimensions**: storage_units_i(32)
- **Attributes**:
  - _FillValue: nan

### storage_units_t_p_i
- **Type**: <class 'str'>
- **Shape**: (32,)
- **Dimensions**: storage_units_t_p_i(32)

### storage_units_t_p
- **Type**: float64
- **Shape**: (56, 32)
- **Dimensions**: snapshots(56), storage_units_t_p_i(32)
- **Attributes**:
  - _FillValue: nan

### storage_units_t_p_dispatch_i
- **Type**: <class 'str'>
- **Shape**: (32,)
- **Dimensions**: storage_units_t_p_dispatch_i(32)

### storage_units_t_p_dispatch
- **Type**: float64
- **Shape**: (56, 32)
- **Dimensions**: snapshots(56), storage_units_t_p_dispatch_i(32)
- **Attributes**:
  - _FillValue: nan

### storage_units_t_p_store_i
- **Type**: <class 'str'>
- **Shape**: (32,)
- **Dimensions**: storage_units_t_p_store_i(32)

### storage_units_t_p_store
- **Type**: float64
- **Shape**: (56, 32)
- **Dimensions**: snapshots(56), storage_units_t_p_store_i(32)
- **Attributes**:
  - _FillValue: nan

### storage_units_t_state_of_charge_i
- **Type**: <class 'str'>
- **Shape**: (32,)
- **Dimensions**: storage_units_t_state_of_charge_i(32)

### storage_units_t_state_of_charge
- **Type**: float64
- **Shape**: (56, 32)
- **Dimensions**: snapshots(56), storage_units_t_state_of_charge_i(32)
- **Attributes**:
  - _FillValue: nan

### storage_units_t_mu_upper_i
- **Type**: <class 'str'>
- **Shape**: (32,)
- **Dimensions**: storage_units_t_mu_upper_i(32)

### storage_units_t_mu_upper
- **Type**: float64
- **Shape**: (56, 32)
- **Dimensions**: snapshots(56), storage_units_t_mu_upper_i(32)
- **Attributes**:
  - _FillValue: nan

### storage_units_t_mu_lower_i
- **Type**: <class 'str'>
- **Shape**: (32,)
- **Dimensions**: storage_units_t_mu_lower_i(32)

### storage_units_t_mu_lower
- **Type**: float64
- **Shape**: (56, 32)
- **Dimensions**: snapshots(56), storage_units_t_mu_lower_i(32)
- **Attributes**:
  - _FillValue: nan

### storage_units_t_mu_energy_balance_i
- **Type**: <class 'str'>
- **Shape**: (32,)
- **Dimensions**: storage_units_t_mu_energy_balance_i(32)

### storage_units_t_mu_energy_balance
- **Type**: float64
- **Shape**: (56, 32)
- **Dimensions**: snapshots(56), storage_units_t_mu_energy_balance_i(32)
- **Attributes**:
  - _FillValue: nan

### links_i
- **Type**: <class 'str'>
- **Shape**: (280,)
- **Dimensions**: links_i(280)

### links_bus0
- **Type**: <class 'str'>
- **Shape**: (280,)
- **Dimensions**: links_i(280)

### links_bus1
- **Type**: <class 'str'>
- **Shape**: (280,)
- **Dimensions**: links_i(280)

### links_p_nom
- **Type**: float64
- **Shape**: (280,)
- **Dimensions**: links_i(280)
- **Attributes**:
  - _FillValue: nan

### links_carrier
- **Type**: <class 'str'>
- **Shape**: (280,)
- **Dimensions**: links_i(280)

### links_underwater_fraction
- **Type**: float64
- **Shape**: (280,)
- **Dimensions**: links_i(280)
- **Attributes**:
  - _FillValue: nan

### links_efficiency
- **Type**: float64
- **Shape**: (280,)
- **Dimensions**: links_i(280)
- **Attributes**:
  - _FillValue: nan

### links_length
- **Type**: float64
- **Shape**: (280,)
- **Dimensions**: links_i(280)
- **Attributes**:
  - _FillValue: nan

### links_capital_cost
- **Type**: float64
- **Shape**: (280,)
- **Dimensions**: links_i(280)
- **Attributes**:
  - _FillValue: nan

### links_p_nom_min
- **Type**: float64
- **Shape**: (280,)
- **Dimensions**: links_i(280)
- **Attributes**:
  - _FillValue: nan

### links_marginal_cost
- **Type**: float64
- **Shape**: (280,)
- **Dimensions**: links_i(280)
- **Attributes**:
  - _FillValue: nan

### links_p_nom_opt
- **Type**: float64
- **Shape**: (280,)
- **Dimensions**: links_i(280)
- **Attributes**:
  - _FillValue: nan

### links_t_p0_i
- **Type**: <class 'str'>
- **Shape**: (280,)
- **Dimensions**: links_t_p0_i(280)

### links_t_p0
- **Type**: float64
- **Shape**: (56, 280)
- **Dimensions**: snapshots(56), links_t_p0_i(280)
- **Attributes**:
  - _FillValue: nan

### links_t_p1_i
- **Type**: <class 'str'>
- **Shape**: (280,)
- **Dimensions**: links_t_p1_i(280)

### links_t_p1
- **Type**: float64
- **Shape**: (56, 280)
- **Dimensions**: snapshots(56), links_t_p1_i(280)
- **Attributes**:
  - _FillValue: nan

### links_t_mu_lower_i
- **Type**: <class 'str'>
- **Shape**: (280,)
- **Dimensions**: links_t_mu_lower_i(280)

### links_t_mu_lower
- **Type**: float64
- **Shape**: (56, 280)
- **Dimensions**: snapshots(56), links_t_mu_lower_i(280)
- **Attributes**:
  - _FillValue: nan

### links_t_mu_upper_i
- **Type**: <class 'str'>
- **Shape**: (275,)
- **Dimensions**: links_t_mu_upper_i(275)

### links_t_mu_upper
- **Type**: float64
- **Shape**: (56, 275)
- **Dimensions**: snapshots(56), links_t_mu_upper_i(275)
- **Attributes**:
  - _FillValue: nan

### carriers_i
- **Type**: <class 'str'>
- **Shape**: (19,)
- **Dimensions**: carriers_i(19)

### carriers_color
- **Type**: <class 'str'>
- **Shape**: (19,)
- **Dimensions**: carriers_i(19)

### carriers_nice_name
- **Type**: <class 'str'>
- **Shape**: (19,)
- **Dimensions**: carriers_i(19)

### carriers_co2_emissions
- **Type**: float64
- **Shape**: (19,)
- **Dimensions**: carriers_i(19)
- **Attributes**:
  - _FillValue: nan

### buses_i
- **Type**: <class 'str'>
- **Shape**: (70,)
- **Dimensions**: buses_i(70)

### buses_Pd
- **Type**: float64
- **Shape**: (70,)
- **Dimensions**: buses_i(70)
- **Attributes**:
  - _FillValue: nan

### buses_v_nom
- **Type**: float64
- **Shape**: (70,)
- **Dimensions**: buses_i(70)
- **Attributes**:
  - _FillValue: nan

### buses_country
- **Type**: <class 'str'>
- **Shape**: (70,)
- **Dimensions**: buses_i(70)

### buses_county
- **Type**: <class 'str'>
- **Shape**: (70,)
- **Dimensions**: buses_i(70)

### buses_reeds_zone
- **Type**: <class 'str'>
- **Shape**: (70,)
- **Dimensions**: buses_i(70)

### buses_reeds_ba
- **Type**: <class 'str'>
- **Shape**: (70,)
- **Dimensions**: buses_i(70)

### buses_interconnect
- **Type**: <class 'str'>
- **Shape**: (70,)
- **Dimensions**: buses_i(70)

### buses_x
- **Type**: float64
- **Shape**: (70,)
- **Dimensions**: buses_i(70)
- **Attributes**:
  - _FillValue: nan

### buses_y
- **Type**: float64
- **Shape**: (70,)
- **Dimensions**: buses_i(70)
- **Attributes**:
  - _FillValue: nan

### buses_nerc_reg
- **Type**: <class 'str'>
- **Shape**: (70,)
- **Dimensions**: buses_i(70)

### buses_trans_reg
- **Type**: <class 'str'>
- **Shape**: (70,)
- **Dimensions**: buses_i(70)

### buses_trans_grp
- **Type**: <class 'str'>
- **Shape**: (70,)
- **Dimensions**: buses_i(70)

### buses_reeds_state
- **Type**: <class 'str'>
- **Shape**: (70,)
- **Dimensions**: buses_i(70)

### buses_rec_trading_zone
- **Type**: <class 'str'>
- **Shape**: (70,)
- **Dimensions**: buses_i(70)

### buses_substation_lv
- **Type**: float64
- **Shape**: (70,)
- **Dimensions**: buses_i(70)
- **Attributes**:
  - _FillValue: nan

### buses_control
- **Type**: <class 'str'>
- **Shape**: (70,)
- **Dimensions**: buses_i(70)

### buses_generator
- **Type**: <class 'str'>
- **Shape**: (70,)
- **Dimensions**: buses_i(70)

### buses_carrier
- **Type**: <class 'str'>
- **Shape**: (70,)
- **Dimensions**: buses_i(70)

### buses_t_p_i
- **Type**: <class 'str'>
- **Shape**: (64,)
- **Dimensions**: buses_t_p_i(64)

### buses_t_p
- **Type**: float64
- **Shape**: (56, 64)
- **Dimensions**: snapshots(56), buses_t_p_i(64)
- **Attributes**:
  - _FillValue: nan

### buses_t_marginal_price_i
- **Type**: <class 'str'>
- **Shape**: (70,)
- **Dimensions**: buses_t_marginal_price_i(70)

### buses_t_marginal_price
- **Type**: float64
- **Shape**: (56, 70)
- **Dimensions**: snapshots(56), buses_t_marginal_price_i(70)
- **Attributes**:
  - _FillValue: nan

### global_constraints_i
- **Type**: <class 'str'>
- **Shape**: (1,)
- **Dimensions**: global_constraints_i(1)

### global_constraints_type
- **Type**: <class 'str'>
- **Shape**: (1,)
- **Dimensions**: global_constraints_i(1)

### global_constraints_carrier_attribute
- **Type**: <class 'str'>
- **Shape**: (1,)
- **Dimensions**: global_constraints_i(1)

### global_constraints_sense
- **Type**: <class 'str'>
- **Shape**: (1,)
- **Dimensions**: global_constraints_i(1)

### global_constraints_constant
- **Type**: float64
- **Shape**: (1,)
- **Dimensions**: global_constraints_i(1)
- **Attributes**:
  - _FillValue: nan

### loads_i
- **Type**: <class 'str'>
- **Shape**: (58,)
- **Dimensions**: loads_i(58)

### loads_bus
- **Type**: <class 'str'>
- **Shape**: (58,)
- **Dimensions**: loads_i(58)

### loads_carrier
- **Type**: <class 'str'>
- **Shape**: (58,)
- **Dimensions**: loads_i(58)

### loads_t_p_set_i
- **Type**: <class 'str'>
- **Shape**: (58,)
- **Dimensions**: loads_t_p_set_i(58)

### loads_t_p_set
- **Type**: float64
- **Shape**: (56, 58)
- **Dimensions**: snapshots(56), loads_t_p_set_i(58)
- **Attributes**:
  - _FillValue: nan

### loads_t_p_i
- **Type**: <class 'str'>
- **Shape**: (58,)
- **Dimensions**: loads_t_p_i(58)

### loads_t_p
- **Type**: float64
- **Shape**: (56, 58)
- **Dimensions**: snapshots(56), loads_t_p_i(58)
- **Attributes**:
  - _FillValue: nan

