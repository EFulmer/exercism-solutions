"""Functions to prevent a nuclear meltdown."""
Numeric = int | float

def is_criticality_balanced(temperature: Numeric, neutrons_emitted: Numeric) -> bool:
    """Verify criticality is balanced.

    Parameters:
        temperature (int or float): The temperature value in kelvin.
        neutrons_emitted (int or float): The number of neutrons emitted per second.

    Returns:
        bool: Is criticality balanced?

    Note:
        A reactor is said to be balanced in criticality if it satisfies the following conditions:
            - The temperature is less than 800 K.
            - The number of neutrons emitted per second is greater than 500.
            - The product of temperature and neutrons emitted per second is less than 500000.

    """
    return is_temperature_safe(temperature) and are_emissions_satisfactory(neutrons_emitted) and are_temp_and_emissions_balanced(temperature, neutrons_emitted)

def is_temperature_safe(temp: Numeric) -> bool:
    return temp < 800


def are_emissions_satisfactory(neutrons_per_second: Numeric) -> bool:
    return neutrons_per_second > 500


def are_temp_and_emissions_balanced(temp: Numeric, neutrons_per_second: Numeric) -> bool:
    return temp * neutrons_per_second < 500_000

def reactor_efficiency(voltage, current, theoretical_max_power):
    """Assess reactor efficiency zone.

    Parameters:
        voltage (int or float): Voltage value.
        current (int or float): Current value.
        theoretical_max_power (int or float): The power level that corresponds to a 100% efficiency.

    Returns:
        str: One of ('green', 'orange', 'red', or 'black').

    Note:
        Efficiency can be grouped into 4 bands:
            1. green -> efficiency of 80% or more,
            2. orange -> efficiency of less than 80% but at least 60%,
            3. red -> efficiency below 60%, but still 30% or more,
            4. black ->  less than 30% efficient.

        The percentage value is calculated as
        (generated power/ theoretical max power)*100
        where generated power = voltage * current
    """
    efficiency = generated_power(voltage, current) / theoretical_max_power
    efficiency_percentage = efficiency * 100
    match efficiency_percentage:
        case percentage if percentage >= 80:
            return "green"
        case percentage if 80 > percentage >= 60:
            return "orange"
        case percentage if 60 > percentage >= 30:
            return "red"
        case _:
            return "black"

def generated_power(voltage: Numeric, current: Numeric) -> Numeric:
    return voltage * current


def fail_safe(temperature: Numeric, neutrons_produced_per_second: Numeric, threshold: Numeric) -> str:
    """Assess and return status code for the reactor.

    Parameters:
        temperature (int or float): The value of the temperature in kelvin.
        neutrons_produced_per_second (int or float): The neutron flux.
        threshold (int or float): The threshold for the category.

    Returns:
        str: One of ('LOW', 'NORMAL', 'DANGER').

    Note:
        1. 'LOW' -> `temperature * neutrons per second` < 90% of `threshold`
        2. 'NORMAL' -> `temperature * neutrons per second` +/- 10% of `threshold`
        3. 'DANGER' -> `temperature * neutrons per second` is not in the above-stated ranges
    """
    criticality = temperature * neutrons_produced_per_second
    bottom_end, top_end = threshold * .9, threshold * 1.1
    if criticality < bottom_end:
        return "LOW"
    if bottom_end <= criticality <= top_end:
        return "NORMAL"
    return "DANGER"