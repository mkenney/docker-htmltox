package chrome

import (
	"app/chrome/protocol"
)

/*
Page - https://chromedevtools.github.io/devtools-protocol/tot/Page/
Actions and events related to the inspected page belong to the page domain.
*/
type Page struct{}

/*
AddScriptToEvaluateOnLoad is eprecated, please use addScriptToEvaluateOnNewDocument instead.
EXPERIMENTAL DEPRECATED
*/
func (Page) AddScriptToEvaluateOnLoad(socket *Socket, params *Page.AddScriptToEvaluateOnLoadParams) error {
	command := &protocol.Command{
		method: "Page.addScriptToEvaluateOnLoad",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
AddScriptToEvaluateOnNewDocument evaluates given script in every frame upon creation (before loading
frame's scripts).
*/
func (Page) AddScriptToEvaluateOnNewDocument(socket *Socket, params *Page.AddScriptToEvaluateOnNewDocumentParams) error {
	command := &protocol.Command{
		method: "Page.addScriptToEvaluateOnNewDocument",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
BringToFront brings page to front (activates tab).
*/
func (Page) BringToFront(socket *Socket) error {
	command := &protocol.Command{
		method: "Page.bringToFront",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
CaptureScreenshot capture a page screenshot.
*/
func (Page) CaptureScreenshot(socket *Socket, params *Page.CaptureScreenshotParams) error {
	command := &protocol.Command{
		method: "Page.CaptureScreenshot",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
CreateIsolatedWorld creates an isolated world for the given frame.
*/
func (Page) CreateIsolatedWorld(socket *Socket, params *Page.CreateIsolatedWorldParams) error {
	command := &protocol.Command{
		method: "Page.createIsolatedWorld",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables page domain notifications.
*/
func (Page) Disable(socket *Socket) error {
	command := &protocol.Command{
		method: "Page.disable",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Ennables page domain notifications.
*/
func (Page) Enable(socket *Socket) error {
	command := &protocol.Command{
		method: "Page.enable",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetAppManifest gets the app manifest.
*/
func (Page) GetAppManifest(socket *Socket, params *Page.GetAppManifestParams) error {
	command := &protocol.Command{
		method: "Page.getAppManifest",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetFrameTree returns present frame tree structure.
*/
func (Page) GetFrameTree(socket *Socket, params *Page.GetFrameTreeParams) error {
	command := &protocol.Command{
		method: "Page.getFrameTree",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetLayoutMetrics returns metrics relating to the layouting of the page, such as viewport
bounds/scale.
*/
func (Page) GetLayoutMetrics(socket *Socket, params *Page.GetLayoutMetricsParams) error {
	command := &protocol.Command{
		method: "Page.getLayoutMetrics",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetNavigationHistory returns navigation history for the current page.
*/
func (Page) GetNavigationHistory(socket *Socket, params *Page.GetNavigationHistoryParams) error {
	command := &protocol.Command{
		method: "Page.getNavigationHistory",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetResourceContent returns content of the given resource. EXPERIMENTAL
*/
func (Page) GetResourceContent(socket *Socket, params *Page.GetResourceContentParams) error {
	command := &protocol.Command{
		method: "Page.getResourceContent",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetResourceTree returns present frame / resource tree structure. EXPERIMENTAL
*/
func (Page) GetResourceTree(socket *Socket, params *Page.GetResourceTreeParams) error {
	command := &protocol.Command{
		method: "Page.getResourceTree",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
HandleJavaScriptDialog accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt,
or onbeforeunload).
*/
func (Page) HandleJavaScriptDialog(socket *Socket, params *Page.HandleJavaScriptDialogParams) error {
	command := &protocol.Command{
		method: "Page.handleJavaScriptDialog",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Navigate navigates current page to the given URL.
*/
func (Page) Navigate(socket *Socket, params *Page.NavigateParams) error {
	command := &protocol.Command{
		method: "Page.navigate",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
NavigateToHistoryEntry navigates current page to the given history entry.
*/
func (Page) NavigateToHistoryEntry(socket *Socket, params *Page.NavigateToHistoryEntryParams) error {
	command := &protocol.Command{
		method: "Page.navigateToHistoryEntry",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
PrintToPDF print page as PDF.
*/
func (Page) PrintToPDF(socket *Socket, params *Page.PrintToPDFParams) error {
	command := &protocol.Command{
		method: "Page.printToPDF",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Reload reloads given page optionally ignoring the cache.
*/
func (Page) Reload(socket *Socket, params *Page.ReloadParams) error {
	command := &protocol.Command{
		method: "Page.reload",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RemoveScriptToEvaluateOnLoad deprecated, please use removeScriptToEvaluateOnNewDocument instead.
EXPERIMENTAL DEPRECATED
*/
func (Page) RemoveScriptToEvaluateOnLoad(socket *Socket, params *Page.RemoveScriptToEvaluateOnLoadParams) error {
	command := &protocol.Command{
		method: "Page.removeScriptToEvaluateOnLoad",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RemoveScriptToEvaluateOnNewDocument removes given script from the list.
*/
func (Page) RemoveScriptToEvaluateOnNewDocument(socket *Socket, params *Page.RemoveScriptToEvaluateOnNewDocumentParams) error {
	command := &protocol.Command{
		method: "Page.removeScriptToEvaluateOnNewDocument",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RequestAppBanner EXPERIMENTAL
*/
func (Page) RequestAppBanner(socket *Socket, params *Page.RequestAppBannerParams) error {
	command := &protocol.Command{
		method: "Page.requestAppBanner",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ScreencastFrameAck acknowledges that a screencast frame has been received by the frontend.
EXPERIMENTAL
*/
func (Page) ScreencastFrameAck(socket *Socket, params *Page.ScreencastFrameAckParams) error {
	command := &protocol.Command{
		method: "Page.screencastFrameAck",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SearchInResource searches for given string in resource content. EXPERIMENTAL
*/
func (Page) SearchInResource(socket *Socket, params *Page.SearchInResourceParams) error {
	command := &protocol.Command{
		method: "Page.searchInResource",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetAdBlockingEnabled enable Chrome's experimental ad filter on all sites. EXPERIMENTAL
*/
func (Page) SetAdBlockingEnabled(socket *Socket, params *Page.SetAdBlockingEnabledParams) error {
	command := &protocol.Command{
		method: "Page.setAdBlockingEnabled",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetAutoAttachToCreatedPages controls whether browser will open a new inspector window for connected
pages. EXPERIMENTAL
*/
func (Page) SetAutoAttachToCreatedPages(socket *Socket, params *Page.SetAutoAttachToCreatedPagesParams) error {
	command := &protocol.Command{
		method: "Page.setAutoAttachToCreatedPages",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetDocumentContent sets given markup as the document's HTML.
*/
func (Page) SetDocumentContent(socket *Socket, params *Page.SetDocumentContentParams) error {
	command := &protocol.Command{
		method: "Page.setDocumentContent",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetDownloadBehavior sets the behavior when downloading a file. EXPERIMENTAL
*/
func (Page) SetDownloadBehavior(socket *Socket, params *Page.SetDownloadBehaviorParams) error {
	command := &protocol.Command{
		method: "Page.setDownloadBehavior",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetLifecycleEventsEnabled controls whether page will emit lifecycle events. EXPERIMENTAL
*/
func (Page) SetLifecycleEventsEnabled(socket *Socket, params *Page.SetLifecycleEventsEnabledParams) error {
	command := &protocol.Command{
		method: "Page.setLifecycleEventsEnabled",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StartScreencast starts sending each frame using the `screencastFrame` event. EXPERIMENTAL
*/
func (Page) StartScreencast(socket *Socket, params *Page.StartScreencastParams) error {
	command := &protocol.Command{
		method: "Page.startScreencast",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopLoading force the page stop all navigations and pending resource fetches.
*/
func (Page) StopLoading(socket *Socket) error {
	command := &protocol.Command{
		method: "Page.stopLoading",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopScreencast stops sending each frame in the `screencastFrame`. EXPERIMENTAL
*/
func (Page) StopScreencast(socket *Socket) error {
	command := &protocol.Command{
		method: "Page.stopScreencast",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}
