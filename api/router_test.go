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
		{"/Albums/testID/Similar", http.StatusInternalServerError, "Not yet implemented"},
		{"/Albums/testID/InstantMix", http.StatusInternalServerError, "Not yet implemented"},

		{"/Artists/testID/Similar", http.StatusInternalServerError, "Not yet implemented"},
		{"/Artists/testName/Images/testType", http.StatusInternalServerError, "Not yet implemented"},
		{"/Artists/testName/Images/testType/testIndex", http.StatusInternalServerError, "Not yet implemented"},
		{"/Artists/testName/InstantMix", http.StatusInternalServerError, "Not yet implemented"},
		{"/Artists/InstantMix", http.StatusInternalServerError, "Not yet implemented"},

		{"/Audio/testID/stream", http.StatusInternalServerError, "Not yet implemented"},
		{"/Audio/testID/stream.m3u8", http.StatusInternalServerError, "Not yet implemented"},
		{"/Audio/testID/hls/segmentId/stream.mp3", http.StatusInternalServerError, "Not yet implemented"},
		{"/Audio/testID/hls1/playlistId/segmentId.ts", http.StatusInternalServerError, "Not yet implemented"},

		{"/Auth/Keys", http.StatusInternalServerError, "Not yet implemented"},

		{"/Branding/Configuration", http.StatusInternalServerError, "Not yet implemented"},
		{"/Branding/Css", http.StatusInternalServerError, "Not yet implemented"},

		{"/Channels", http.StatusInternalServerError, "Not yet implemented"},
		{"/Channels/Features", http.StatusInternalServerError, "Not yet implemented"},
		{"/Channels/Folder", http.StatusInternalServerError, "Not yet implemented"},
		{"/Channels/testID/Features", http.StatusInternalServerError, "Not yet implemented"},
		{"/Channels/testID/Items", http.StatusInternalServerError, "Not yet implemented"},
		{"/Channels/testID/Latest", http.StatusInternalServerError, "Not yet implemented"},

		{"/Collections", http.StatusInternalServerError, "Not yet implemented"},

		{"/Connect/Pending", http.StatusInternalServerError, "Not yet implemented"},
		{"/Connect/Exchange", http.StatusInternalServerError, "Not yet implemented"},

		{"/Devices", http.StatusInternalServerError, "Not yet implemented"},
		{"/Devices/CameraUploads", http.StatusInternalServerError, "Not yet implemented"},
		{"/Devices/Info", http.StatusInternalServerError, "Not yet implemented"},
		{"/Devices/Capabilities", http.StatusInternalServerError, "Not yet implemented"},

		{"/DisplayPreferences/testId", http.StatusInternalServerError, "Not yet implemented"},

		{"/Dlna/testId/description", http.StatusInternalServerError, "Not yet implemented"},
		{"/Dlna/testId/contentdirectory/contentdirectory.xml", http.StatusInternalServerError, "Not yet implemented"},
		{"/Dlna/testId/connectionmanager/connectionmanager.xml", http.StatusInternalServerError, "Not yet implemented"},
		{"/Dlna/testId/mediareceiverregistrar/mediareceiverregistrar.xml", http.StatusInternalServerError, "Not yet implemented"},
		{"/Dlna/testId/icons/testfilename", http.StatusInternalServerError, "Not yet implemented"},
		{"/Dlna/ProfileInfos", http.StatusInternalServerError, "Not yet implemented"},
		{"/Dlna/Profiles/Default", http.StatusInternalServerError, "Not yet implemented"},
		{"/Dlna/Profiles/testId", http.StatusInternalServerError, "Not yet implemented"},

		{"/Environment/DirectoryContents", http.StatusInternalServerError, "Not yet implemented"},
		{"/Environment/NetworkShares", http.StatusInternalServerError, "Not yet implemented"},
		{"/Environment/NetworkDevices", http.StatusInternalServerError, "Not yet implemented"},
		{"/Environment/Drives", http.StatusInternalServerError, "Not yet implemented"},
		{"/Environment/ParentPath", http.StatusInternalServerError, "Not yet implemented"},

		{"/Games/testId/Similar", http.StatusInternalServerError, "Not yet implemented"},
		{"/Games/SystemSummaries", http.StatusInternalServerError, "Not yet implemented"},
		{"/Games/PlayerIndex", http.StatusInternalServerError, "Not yet implemented"},

		{"/GameGenres", http.StatusInternalServerError, "Not yet implemented"},
		{"/GameGenres/testName", http.StatusInternalServerError, "Not yet implemented"},
		{"/GameGenres/testName/Images/testType/testIndex", http.StatusInternalServerError, "Not yet implemented"},
		{"/GameGenres/testName", http.StatusInternalServerError, "Not yet implemented"},

		{"/Genres", http.StatusInternalServerError, "Not yet implemented"},
		{"/Genres/testName", http.StatusInternalServerError, "Not yet implemented"},
		{"/Genres/testName/Images/testType/testIndex", http.StatusInternalServerError, "Not yet implemented"},
		{"/Genres/testName", http.StatusInternalServerError, "Not yet implemented"},

		{"/Images/General", http.StatusInternalServerError, "Not yet implemented"},
		{"/Images/Ratings", http.StatusInternalServerError, "Not yet implemented"},
		{"/Images/MediaInfo", http.StatusInternalServerError, "Not yet implemented"},
		{"/Images/General/testName/testType", http.StatusInternalServerError, "Not yet implemented"},
		{"/Images/Ratings/testTheme/testName", http.StatusInternalServerError, "Not yet implemented"},
		{"/Images/MediaInfo/testTheme/testName", http.StatusInternalServerError, "Not yet implemented"},
		{"/Images/Remote", http.StatusInternalServerError, "Not yet implemented"},

		{"/Items/Counts", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/Filters", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testID/Ancestors", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testID/CriticReviews", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testID/Download", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testID/File", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testID/Similar", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testID/ThemeMedia", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testID/ThemeSongs", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testID/ThemeVideos", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/YearIndex", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testId/Images", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testId/Images/testType", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testId/Images/testType/testIndex", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testId/Images/testType/testIndex/tag/format/maxwidth/maxheight/percentplayed/unplayedcount", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testId/InstantMix", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testId/ExternalIdInfos", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/RemoteSearch/Image", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testId/MetadataEditor", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testId/PlaybackInfo", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testId/RemoteImages", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testId/RemoteImages/Providers", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testId/RemoteSearch/Subtitles/Providers", http.StatusInternalServerError, "Not yet implemented"},
		{"/Items/testId/RemoteSearch/Subtitles/testlang", http.StatusInternalServerError, "Not yet implemented"},

		{"/Library/FileOrganization", http.StatusInternalServerError, "Not yet implemented"},
		{"/Library/FileOrganization/SmartMatches", http.StatusInternalServerError, "Not yet implemented"},
		{"/Library/MediaFolders", http.StatusInternalServerError, "Not yet implemented"},
		{"/Library/PhysicalPaths", http.StatusInternalServerError, "Not yet implemented"},
		{"/Library/VirtualFolders", http.StatusInternalServerError, "Not yet implemented"},

		{"/LiveTv/Info", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/Channels", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/Channels/testId", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/Recordings", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/Recordings/Groups", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/Recordings/testId", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/Timers/testId", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/Timers/Defaults", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/Timers", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/Programs", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/Programs/Recommended", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/Programs/testId", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/SeriesTimers/testId", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/SeriesTimers", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/Recordings/Groups/testId", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/GuideInfo", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/Folder", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/ListingProviders/Lineups", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/ListingProviders/SchedulesDirect/Countries", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/Registration", http.StatusInternalServerError, "Not yet implemented"},
		{"/LiveTv/TunerHosts/Satip/IniMappings", http.StatusInternalServerError, "Not yet implemented"},

		{"/Localization/Cultures", http.StatusInternalServerError, "Not yet implemented"},
		{"/Localization/Countries", http.StatusInternalServerError, "Not yet implemented"},
		{"/Localization/ParentalRatings", http.StatusInternalServerError, "Not yet implemented"},
		{"/Localization/Options", http.StatusInternalServerError, "Not yet implemented"},

		{"/Movies/testId/Similar", http.StatusInternalServerError, "Not yet implemented"},
		{"/Movies/Recommendations", http.StatusInternalServerError, "Not yet implemented"},

		{"/MusicGenres/testName/Images/testType", http.StatusInternalServerError, "Not yet implemented"},
		{"/MusicGenres/testName/Images/testType/testIndex", http.StatusInternalServerError, "Not yet implemented"},
		{"/MusicGenres/testName/InstantMix", http.StatusInternalServerError, "Not yet implemented"},
		{"/MusicGenres", http.StatusInternalServerError, "Not yet implemented"},
		{"/MusicGenres/testName", http.StatusInternalServerError, "Not yet implemented"},

		{"/News/Product", http.StatusInternalServerError, "Not yet implemented"},

		{"/Notifications/testId", http.StatusInternalServerError, "Not yet implemented"},
		{"/Notifications/testId/Summary", http.StatusInternalServerError, "Not yet implemented"},
		{"/Notifications/Types", http.StatusInternalServerError, "Not yet implemented"},
		{"/Notifications/Services", http.StatusInternalServerError, "Not yet implemented"},

		{"/Packages/Reviews/testId", http.StatusInternalServerError, "Not yet implemented"},

		{"/Persons/testName/Images/testType", http.StatusInternalServerError, "Not yet implemented"},
		{"/Persons/testName/Images/testType/testIndex", http.StatusInternalServerError, "Not yet implemented"},
		{"/Persons", http.StatusInternalServerError, "Not yet implemented"},
		{"/Persons/testName", http.StatusInternalServerError, "Not yet implemented"},

		{"/Playback/BitrateTest", http.StatusInternalServerError, "Not yet implemented"},

		{"/Playlists/testId/Items", http.StatusInternalServerError, "Not yet implemented"},
		{"/Playlists/testId/InstantMix", http.StatusInternalServerError, "Not yet implemented"},

		{"/Providers/Chapters", http.StatusInternalServerError, "Not yet implemented"},
		{"/Providers/Subtitles/Subtitles/testId", http.StatusInternalServerError, "Not yet implemented"},

		{"/ScheduledTasks/testId", http.StatusInternalServerError, "Not yet implemented"},
		{"/ScheduledTasks", http.StatusInternalServerError, "Not yet implemented"},

		{"/Search/Hints", http.StatusInternalServerError, "Not yet implemented"},

		{"/Sessions", http.StatusInternalServerError, "Not yet implemented"},

		{"/Shows/NextUp", http.StatusInternalServerError, "Not yet implemented"},
		{"/Shows/Upcoming", http.StatusInternalServerError, "Not yet implemented"},
		{"/Shows/testId/Similar", http.StatusInternalServerError, "Not yet implemented"},
		{"/Shows/testId/Episodes", http.StatusInternalServerError, "Not yet implemented"},
		{"/Shows/testId/Seasons", http.StatusInternalServerError, "Not yet implemented"},

		{"/Social/Shares/testId", http.StatusInternalServerError, "Not yet implemented"},
		{"/Social/Shares/Public/testId", http.StatusInternalServerError, "Not yet implemented"},
		{"/Social/Shares/Public/testId/Images", http.StatusInternalServerError, "Not yet implemented"},
		{"/Social/Shares/Public/testId/Item", http.StatusInternalServerError, "Not yet implemented"},

		{"/Songs/testId/InstantMix", http.StatusInternalServerError, "Not yet implemented"},

		{"/Startup/Complete", http.StatusInternalServerError, "Not yet implemented"},
		{"/Startup/Configuration", http.StatusInternalServerError, "Not yet implemented"},
		{"/Startup/User", http.StatusInternalServerError, "Not yet implemented"},

		{"/Studios/testName/Images/testType", http.StatusInternalServerError, "Not yet implemented"},
		{"/Studios/testName/Images/testType/testIndex", http.StatusInternalServerError, "Not yet implemented"},
		{"/Studios", http.StatusInternalServerError, "Not yet implemented"},
		{"/Studios/testName", http.StatusInternalServerError, "Not yet implemented"},

		{"/Sync/Jobs/testId", http.StatusInternalServerError, "Not yet implemented"},
		{"/Sync/JobItems", http.StatusInternalServerError, "Not yet implemented"},
		{"/Sync/Jobs", http.StatusInternalServerError, "Not yet implemented"},
		{"/Sync/Targets", http.StatusInternalServerError, "Not yet implemented"},
		{"/Sync/Options", http.StatusInternalServerError, "Not yet implemented"},
		{"/Sync/JobItems/testId/File", http.StatusInternalServerError, "Not yet implemented"},
		{"/Sync/JobItems/testId/AdditionalFiles", http.StatusInternalServerError, "Not yet implemented"},
		{"/Sync/Items/Ready", http.StatusInternalServerError, "Not yet implemented"},

		{"/System/ActivityLog/Entries", http.StatusInternalServerError, "Not yet implemented"},
		{"/System/Configuration/testKey", http.StatusInternalServerError, "Not yet implemented"},
		{"/System/Configuration/MetadataOptions/Default", http.StatusInternalServerError, "Not yet implemented"},
		{"/System/Configuration/MetadataPlugins", http.StatusInternalServerError, "Not yet implemented"},
		{"/System/Info", http.StatusInternalServerError, "Not yet implemented"},
		{"/System/Info/Public", http.StatusInternalServerError, "Not yet implemented"},
		{"/System/Logs", http.StatusInternalServerError, "Not yet implemented"},
		{"/System/Logs/Log", http.StatusInternalServerError, "Not yet implemented"},
		{"/System/Endpoint", http.StatusInternalServerError, "Not yet implemented"},

		{"/Trailers", http.StatusInternalServerError, "Not yet implemented"},
		{"/Trailers/testId/Similar", http.StatusInternalServerError, "Not yet implemented"},

		{"/Users", http.StatusInternalServerError, "Not yet implemented"},
		{"/Users/Public", http.StatusInternalServerError, "Not yet implemented"},
		{"/Users/testId", http.StatusInternalServerError, "Not yet implemented"},
		{"/Users/testId/Offline", http.StatusInternalServerError, "Not yet implemented"},
		{"/Users/testId/Views", http.StatusInternalServerError, "Not yet implemented"},
		{"/Users/testId/SpecialViewOptions", http.StatusInternalServerError, "Not yet implemented"},
		{"/Users/testId/GroupingOptions", http.StatusInternalServerError, "Not yet implemented"},
		{"/Users/testId/Images/testType", http.StatusInternalServerError, "Not yet implemented"},
		{"/Users/testId/Images/testType/testIndex", http.StatusInternalServerError, "Not yet implemented"},
		{"/Users/testUid/Items/testId", http.StatusInternalServerError, "Not yet implemented"},
		{"/Users/testUid/Items/Root", http.StatusInternalServerError, "Not yet implemented"},
		{"/Users/testUid/Items/testId/Intros", http.StatusInternalServerError, "Not yet implemented"},
		{"/Users/testUid/Items/testId/LocalTrailers", http.StatusInternalServerError, "Not yet implemented"},
		{"/Users/testUid/Items/testId/SpecialFeatures", http.StatusInternalServerError, "Not yet implemented"},
		{"/Users/testUid/Items/Latest", http.StatusInternalServerError, "Not yet implemented"},

		{"/Videos/testId/master.m3u8", http.StatusInternalServerError, "Not yet implemented"},
		{"/Videos/testID/hls/playlistId/teststream.m3u8", http.StatusInternalServerError, "Not yet implemented"},
		{"/Videos/testID/hls/playlistId/segmentId.ts", http.StatusInternalServerError, "Not yet implemented"},
		{"/Videos/testID/hls1/playlistId/segmentId.ts", http.StatusInternalServerError, "Not yet implemented"},
		{"/Videos/testId/master.mpd", http.StatusInternalServerError, "Not yet implemented"},
		{"/Videos/testId/dash/testrepId/testSegmentId.m4s", http.StatusInternalServerError, "Not yet implemented"},
		{"/Videos/testId/Subtitles/testIndex", http.StatusInternalServerError, "Not yet implemented"},
		{"/Videos/testId/mediaSrc/Subtitles/testIndex/streamFormat", http.StatusInternalServerError, "Not yet implemented"},
		{"/Videos/testId/mediaSrc/Subtitles/testIndex/startPos/streamFormat", http.StatusInternalServerError, "Not yet implemented"},
		{"/Videos/testId/mediaSrc/Subtitles/testIndex/subtitles.m3u8", http.StatusInternalServerError, "Not yet implemented"},
		{"/Videos/testId/live.m3u8", http.StatusInternalServerError, "Not yet implemented"},
		{"/Videos/testId/stream.mkv", http.StatusInternalServerError, "Not yet implemented"},
		{"/Videos/testId/AdditionalParts", http.StatusInternalServerError, "Not yet implemented"},

		{"/Years", http.StatusInternalServerError, "Not yet implemented"},
		{"/Years/testYear", http.StatusInternalServerError, "Not yet implemented"},
		{"/Years/testYear/Images/testType", http.StatusInternalServerError, "Not yet implemented"},
		{"/Years/testYear/Images/testType/testIndex", http.StatusInternalServerError, "Not yet implemented"},
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
