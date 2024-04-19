package main

import (
	"fmt"

	"github.com/distatus/battery"
	mp "github.com/mackerelio/go-mackerel-plugin"
)

func main() {
	plugin := new(Plugin)

	mackerelPlugin := mp.NewMackerelPlugin(plugin)
	mackerelPlugin.Run()
}

type Plugin struct{}

var _ mp.PluginWithPrefix = new(Plugin)

func (p *Plugin) FetchMetrics() (map[string]float64, error) {
	batteries, err := battery.GetAll()
	if err != nil {
		return nil, err
	}
	metrics := make(map[string]float64, 3*len(batteries))
	for i, b := range batteries {
		metrics[fmt.Sprintf("capacity.%d.current", i)] = b.Current
		metrics[fmt.Sprintf("capacity.%d.design", i)] = b.Design
		metrics[fmt.Sprintf("capacity.%d.full", i)] = b.Full
		metrics[fmt.Sprintf("usage_rate.%d.percentage", i)] = b.Current / b.Full * 100
	}

	return metrics, nil
}

func (p *Plugin) GraphDefinition() map[string]mp.Graphs {
	return map[string]mp.Graphs{
		"capacity.#": {
			Label: "Battery Capacity [mWh]",
			Unit:  mp.UnitFloat,
			Metrics: []mp.Metrics{
				{Name: "current", Label: "Current Capacity"},
				{Name: "design", Label: "Design Capacity"},
				{Name: "full", Label: "Full Capacity"},
			},
		},
		"usage_rate.#": {
			Label: "Battery Usage Rate",
			Unit:  mp.UnitPercentage,
			Metrics: []mp.Metrics{
				{Name: "percentage", Label: "Battery Usage Rate"},
			},
		},
	}
}

func (p *Plugin) MetricKeyPrefix() string {
	return "battery"
}
