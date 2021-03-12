package main

import (
	"fmt"

	"github.com/getlantern/systray"
)

func main() {
	onExit := func() {}
	// Should be called at the very beginning of main().
	systray.RunWithAppWindow("A15 Links", 1280, 960, onReady, onExit)
}

func onReady() {
	systray.SetTitle("A15-Links")
	systray.SetTooltip("A15 Links")

	google := systray.AddMenuItem("Google", "google search")
	jsonFormatter := systray.AddMenuItem("JSON Formatter", "beautify json")
	jsonDiff := systray.AddMenuItem("JSON Diff Tool", "compare 2 json docs")
	epoch := systray.AddMenuItem("Epoch Converter", "Epoch Converter")
	systray.AddSeparator()

	github := systray.AddMenuItem("URBN Github", "URBN Github")
	ecommInventory := systray.AddMenuItem("ecomm inventory", "ecomm inventory")
	redteamSprint := systray.AddMenuItem("JIRA Redteam Sprint", "JIRA Redteam Sprint")
	pagerDutyRunbooks := systray.AddMenuItem("Pager Duty Runbooks", "Pager Duty Runbooks")
	elkphl := systray.AddMenuItem("ELk PHL", "ELk PHL")
	elkrno := systray.AddMenuItem("ELk RNO", "ELk RNO")

	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	for {
		select {
		case <-google.ClickedCh:
			systray.ShowAppWindow("https://www.google.com")
		case <-jsonFormatter.ClickedCh:
			systray.ShowAppWindow("https://jsonformatter.curiousconcept.com/")
		case <-jsonDiff.ClickedCh:
			systray.ShowAppWindow("http://www.jsondiff.com/")
		case <-epoch.ClickedCh:
			systray.ShowAppWindow("https://www.epochconverter.com/")

		case <-github.ClickedCh:
			systray.ShowAppWindow("https://github.com/urbn")
		case <-ecommInventory.ClickedCh:
			systray.ShowAppWindow("https://gitlab.urbn.com/urbn/ecomm-inventory")
		case <-redteamSprint.ClickedCh:
			systray.ShowAppWindow("https://urbnit.atlassian.net/secure/RapidBoard.jspa?rapidView=338")
		case <-pagerDutyRunbooks.ClickedCh:
			systray.ShowAppWindow("https://urbnit.atlassian.net/wiki/spaces/A15/pages/186896412/Pager+Duty+Alerts+-+Run+Book+Summary")
		case <-elkphl.ClickedCh:
			systray.ShowAppWindow("https://ecomm-elk.urbn.com/app/kibana#/home?_g=()")
		case <-elkrno.ClickedCh:
			fmt.Println("elkrno")
			systray.ShowAppWindow("https://p1inelk101.rno.ecomm.urbanout.com/app/kibana#/home?_g=()")

		case <-mQuit.ClickedCh:
			systray.Quit()
			return
		}
	}
}
