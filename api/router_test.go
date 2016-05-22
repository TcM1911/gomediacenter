package api_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tcm1911/gomediacenter/api"
)

var SERVER_URL string

func init() {
	handler := api.NewAPIRouter()
	server := httptest.NewServer(handler)
	SERVER_URL = server.URL
}

func TestAPIGetRoutes(t *testing.T) {
	tests := []struct {
		url  string
		code int
		body string
	}{
		{"/Albums/testID/Similar", http.StatusNotImplemented, "Not yet implemented"},
		{"/Albums/testID/InstantMix", http.StatusNotImplemented, "Not yet implemented"},

		{"/Artists/testID/Similar", http.StatusNotImplemented, "Not yet implemented"},
		{"/Artists/testName/Images/testType", http.StatusNotImplemented, "Not yet implemented"},
		{"/Artists/testName/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/Artists/testName/InstantMix", http.StatusNotImplemented, "Not yet implemented"},
		{"/Artists/InstantMix", http.StatusNotImplemented, "Not yet implemented"},

		{"/Audio/testID/stream", http.StatusNotImplemented, "Not yet implemented"},
		{"/Audio/testID/stream.m3u8", http.StatusNotImplemented, "Not yet implemented"},
		{"/Audio/testID/hls/segmentId/stream.mp3", http.StatusNotImplemented, "Not yet implemented"},
		{"/Audio/testID/hls1/playlistId/segmentId.ts", http.StatusNotImplemented, "Not yet implemented"},

		{"/Auth/Keys", http.StatusNotImplemented, "Not yet implemented"},

		{"/Branding/Configuration", http.StatusNotImplemented, "Not yet implemented"},
		{"/Branding/Css", http.StatusNotImplemented, "Not yet implemented"},

		{"/Channels", http.StatusNotImplemented, "Not yet implemented"},
		{"/Channels/Features", http.StatusNotImplemented, "Not yet implemented"},
		{"/Channels/Folder", http.StatusNotImplemented, "Not yet implemented"},
		{"/Channels/testID/Features", http.StatusNotImplemented, "Not yet implemented"},
		{"/Channels/testID/Items", http.StatusNotImplemented, "Not yet implemented"},
		{"/Channels/testID/Latest", http.StatusNotImplemented, "Not yet implemented"},

		{"/Collections", http.StatusNotImplemented, "Not yet implemented"},

		{"/Connect/Pending", http.StatusNotImplemented, "Not yet implemented"},
		{"/Connect/Exchange", http.StatusNotImplemented, "Not yet implemented"},

		{"/Devices", http.StatusNotImplemented, "Not yet implemented"},
		{"/Devices/CameraUploads", http.StatusNotImplemented, "Not yet implemented"},
		{"/Devices/Info", http.StatusNotImplemented, "Not yet implemented"},
		{"/Devices/Capabilities", http.StatusNotImplemented, "Not yet implemented"},

		{"/DisplayPreferences/testId", http.StatusNotImplemented, "Not yet implemented"},

		{"/Dlna/testId/description", http.StatusNotImplemented, "Not yet implemented"},
		{"/Dlna/testId/contentdirectory/contentdirectory.xml", http.StatusNotImplemented, "Not yet implemented"},
		{"/Dlna/testId/connectionmanager/connectionmanager.xml", http.StatusNotImplemented, "Not yet implemented"},
		{"/Dlna/testId/mediareceiverregistrar/mediareceiverregistrar.xml", http.StatusNotImplemented, "Not yet implemented"},
		{"/Dlna/testId/icons/testfilename", http.StatusNotImplemented, "Not yet implemented"},
		{"/Dlna/ProfileInfos", http.StatusNotImplemented, "Not yet implemented"},
		{"/Dlna/Profiles/Default", http.StatusNotImplemented, "Not yet implemented"},
		{"/Dlna/Profiles/testId", http.StatusNotImplemented, "Not yet implemented"},

		{"/Environment/DirectoryContents", http.StatusNotImplemented, "Not yet implemented"},
		{"/Environment/NetworkShares", http.StatusNotImplemented, "Not yet implemented"},
		{"/Environment/NetworkDevices", http.StatusNotImplemented, "Not yet implemented"},
		{"/Environment/Drives", http.StatusNotImplemented, "Not yet implemented"},
		{"/Environment/ParentPath", http.StatusNotImplemented, "Not yet implemented"},

		{"/Games/testId/Similar", http.StatusNotImplemented, "Not yet implemented"},
		{"/Games/SystemSummaries", http.StatusNotImplemented, "Not yet implemented"},
		{"/Games/PlayerIndex", http.StatusNotImplemented, "Not yet implemented"},

		{"/GameGenres", http.StatusNotImplemented, "Not yet implemented"},
		{"/GameGenres/testName", http.StatusNotImplemented, "Not yet implemented"},
		{"/GameGenres/testName/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/GameGenres/testName", http.StatusNotImplemented, "Not yet implemented"},

		{"/Genres", http.StatusNotImplemented, "Not yet implemented"},
		{"/Genres/testName", http.StatusNotImplemented, "Not yet implemented"},
		{"/Genres/testName/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/Genres/testName", http.StatusNotImplemented, "Not yet implemented"},

		{"/Images/General", http.StatusNotImplemented, "Not yet implemented"},
		{"/Images/Ratings", http.StatusNotImplemented, "Not yet implemented"},
		{"/Images/MediaInfo", http.StatusNotImplemented, "Not yet implemented"},
		{"/Images/General/testName/testType", http.StatusNotImplemented, "Not yet implemented"},
		{"/Images/Ratings/testTheme/testName", http.StatusNotImplemented, "Not yet implemented"},
		{"/Images/MediaInfo/testTheme/testName", http.StatusNotImplemented, "Not yet implemented"},
		{"/Images/Remote", http.StatusNotImplemented, "Not yet implemented"},

		{"/Items/Counts", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/Filters", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testID/Ancestors", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testID/CriticReviews", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testID/Download", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testID/File", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testID/Similar", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testID/ThemeMedia", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testID/ThemeSongs", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testID/ThemeVideos", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/YearIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/Images", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/Images/testType", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/Images/testType/testIndex/tag/format/maxwidth/maxheight/percentplayed/unplayedcount", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/InstantMix", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/ExternalIdInfos", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/RemoteSearch/Image", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/MetadataEditor", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/PlaybackInfo", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/RemoteImages", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/RemoteImages/Providers", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/RemoteSearch/Subtitles/Providers", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/RemoteSearch/Subtitles/testlang", http.StatusNotImplemented, "Not yet implemented"},

		{"/Library/FileOrganization", http.StatusNotImplemented, "Not yet implemented"},
		{"/Library/FileOrganization/SmartMatches", http.StatusNotImplemented, "Not yet implemented"},
		{"/Library/MediaFolders", http.StatusNotImplemented, "Not yet implemented"},
		{"/Library/PhysicalPaths", http.StatusNotImplemented, "Not yet implemented"},
		{"/Library/VirtualFolders", http.StatusNotImplemented, "Not yet implemented"},

		{"/LiveTv/Info", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Channels", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Channels/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Recordings", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Recordings/Groups", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Recordings/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Timers/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Timers/Defaults", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Timers", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Programs", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Programs/Recommended", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Programs/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/SeriesTimers/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/SeriesTimers", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Recordings/Groups/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/GuideInfo", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Folder", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/ListingProviders/Lineups", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/ListingProviders/SchedulesDirect/Countries", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Registration", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/TunerHosts/Satip/IniMappings", http.StatusNotImplemented, "Not yet implemented"},

		{"/Localization/Cultures", http.StatusNotImplemented, "Not yet implemented"},
		{"/Localization/Countries", http.StatusNotImplemented, "Not yet implemented"},
		{"/Localization/ParentalRatings", http.StatusNotImplemented, "Not yet implemented"},
		{"/Localization/Options", http.StatusNotImplemented, "Not yet implemented"},

		{"/Movies/testId/Similar", http.StatusNotImplemented, "Not yet implemented"},
		{"/Movies/Recommendations", http.StatusNotImplemented, "Not yet implemented"},

		{"/MusicGenres/testName/Images/testType", http.StatusNotImplemented, "Not yet implemented"},
		{"/MusicGenres/testName/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/MusicGenres/testName/InstantMix", http.StatusNotImplemented, "Not yet implemented"},
		{"/MusicGenres", http.StatusNotImplemented, "Not yet implemented"},
		{"/MusicGenres/testName", http.StatusNotImplemented, "Not yet implemented"},

		{"/News/Product", http.StatusNotImplemented, "Not yet implemented"},

		{"/Notifications/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/Notifications/testId/Summary", http.StatusNotImplemented, "Not yet implemented"},
		{"/Notifications/Types", http.StatusNotImplemented, "Not yet implemented"},
		{"/Notifications/Services", http.StatusNotImplemented, "Not yet implemented"},

		{"/Packages/Reviews/testId", http.StatusNotImplemented, "Not yet implemented"},

		{"/Persons/testName/Images/testType", http.StatusNotImplemented, "Not yet implemented"},
		{"/Persons/testName/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/Persons", http.StatusNotImplemented, "Not yet implemented"},
		{"/Persons/testName", http.StatusNotImplemented, "Not yet implemented"},

		{"/Playback/BitrateTest", http.StatusNotImplemented, "Not yet implemented"},

		{"/Playlists/testId/Items", http.StatusNotImplemented, "Not yet implemented"},
		{"/Playlists/testId/InstantMix", http.StatusNotImplemented, "Not yet implemented"},

		{"/Providers/Chapters", http.StatusNotImplemented, "Not yet implemented"},
		{"/Providers/Subtitles/Subtitles/testId", http.StatusNotImplemented, "Not yet implemented"},

		{"/ScheduledTasks/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/ScheduledTasks", http.StatusNotImplemented, "Not yet implemented"},

		{"/Search/Hints", http.StatusNotImplemented, "Not yet implemented"},

		{"/Sessions", http.StatusNotImplemented, "Not yet implemented"},

		{"/Shows/NextUp", http.StatusNotImplemented, "Not yet implemented"},
		{"/Shows/Upcoming", http.StatusNotImplemented, "Not yet implemented"},
		{"/Shows/testId/Similar", http.StatusNotImplemented, "Not yet implemented"},
		{"/Shows/testId/Episodes", http.StatusNotImplemented, "Not yet implemented"},
		{"/Shows/testId/Seasons", http.StatusNotImplemented, "Not yet implemented"},

		{"/Social/Shares/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/Social/Shares/Public/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/Social/Shares/Public/testId/Images", http.StatusNotImplemented, "Not yet implemented"},
		{"/Social/Shares/Public/testId/Item", http.StatusNotImplemented, "Not yet implemented"},

		{"/Songs/testId/InstantMix", http.StatusNotImplemented, "Not yet implemented"},

		{"/Startup/Complete", http.StatusNotImplemented, "Not yet implemented"},
		{"/Startup/Configuration", http.StatusNotImplemented, "Not yet implemented"},
		{"/Startup/User", http.StatusNotImplemented, "Not yet implemented"},

		{"/Studios/testName/Images/testType", http.StatusNotImplemented, "Not yet implemented"},
		{"/Studios/testName/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/Studios", http.StatusNotImplemented, "Not yet implemented"},
		{"/Studios/testName", http.StatusNotImplemented, "Not yet implemented"},

		{"/Sync/Jobs/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/Sync/JobItems", http.StatusNotImplemented, "Not yet implemented"},
		{"/Sync/Jobs", http.StatusNotImplemented, "Not yet implemented"},
		{"/Sync/Targets", http.StatusNotImplemented, "Not yet implemented"},
		{"/Sync/Options", http.StatusNotImplemented, "Not yet implemented"},
		{"/Sync/JobItems/testId/File", http.StatusNotImplemented, "Not yet implemented"},
		{"/Sync/JobItems/testId/AdditionalFiles", http.StatusNotImplemented, "Not yet implemented"},
		{"/Sync/Items/Ready", http.StatusNotImplemented, "Not yet implemented"},

		{"/System/ActivityLog/Entries", http.StatusNotImplemented, "Not yet implemented"},
		{"/System/Configuration/testKey", http.StatusNotImplemented, "Not yet implemented"},
		{"/System/Configuration/MetadataOptions/Default", http.StatusNotImplemented, "Not yet implemented"},
		{"/System/Configuration/MetadataPlugins", http.StatusNotImplemented, "Not yet implemented"},
		{"/System/Info", http.StatusNotImplemented, "Not yet implemented"},
		{"/System/Info/Public", http.StatusNotImplemented, "Not yet implemented"},
		{"/System/Logs", http.StatusNotImplemented, "Not yet implemented"},
		{"/System/Logs/Log", http.StatusNotImplemented, "Not yet implemented"},
		{"/System/Endpoint", http.StatusNotImplemented, "Not yet implemented"},

		{"/Trailers", http.StatusNotImplemented, "Not yet implemented"},
		{"/Trailers/testId/Similar", http.StatusNotImplemented, "Not yet implemented"},

		{"/Users", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/Public", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testId/Offline", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testId/Views", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testId/SpecialViewOptions", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testId/GroupingOptions", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testId/Images/testType", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testId/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testUid/Items/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testUid/Items/Root", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testUid/Items/testId/Intros", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testUid/Items/testId/LocalTrailers", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testUid/Items/testId/SpecialFeatures", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testUid/Items/Latest", http.StatusNotImplemented, "Not yet implemented"},

		{"/Videos/testId/master.m3u8", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testID/hls/playlistId/teststream.m3u8", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testID/hls/playlistId/segmentId.ts", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testID/hls1/playlistId/segmentId.ts", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/master.mpd", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/dash/testrepId/testSegmentId.m4s", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/Subtitles/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/mediaSrc/Subtitles/testIndex/streamFormat", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/mediaSrc/Subtitles/testIndex/startPos/streamFormat", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/mediaSrc/Subtitles/testIndex/subtitles.m3u8", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/live.m3u8", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/stream.mkv", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/AdditionalParts", http.StatusNotImplemented, "Not yet implemented"},

		{"/Years", http.StatusNotImplemented, "Not yet implemented"},
		{"/Years/testYear", http.StatusNotImplemented, "Not yet implemented"},
		{"/Years/testYear/Images/testType", http.StatusNotImplemented, "Not yet implemented"},
		{"/Years/testYear/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
	}

	for _, test := range tests {
		url := SERVER_URL + test.url

		fmt.Printf("Testing GET request to %s\n", url)
		resp, _ := http.Get(url)

		reader := resp.Body
		body, _ := ioutil.ReadAll(reader)
		reader.Close()

		if resp.StatusCode != test.code {
			t.Errorf("Wrong status code for %s. Got: %d, expected %d", test.url, resp.StatusCode, test.code)
		}

		if string(body) != test.body {
			t.Errorf("Wrong reesponse body for %s. Got: %s, expected %s", test.url, body, test.body)
		}
	}

}
