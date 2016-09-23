package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tcm1911/gomediacenter"
)

// This is used for routers that are not yet implemented.
func notYetImplemented(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusNotImplemented)
	rw.Write([]byte("Not yet implemented"))
}

// NewAPIRouter returns a new api router.
func NewAPIRouter(userManager *gomediacenter.UserManager) *(mux.Router) {
	router := mux.NewRouter()

	// /Albums router
	newAlbumsRounter(router)

	// /Artists
	newArtistsRouter(router)

	// /Audio
	newAudioRouter(router)

	// /Auth
	newAuthRouter(router)

	// /Branding
	newBrandingRouter(router)

	// /Channels
	newChannelsRouter(router)

	// /Collections
	newCollectionsRouter(router)

	// /Connect
	newConnectRouter(router)

	// /Devices
	newDevicesRouter(router)

	// /DisplayPreferences
	newDisplayPreferenceRouter(router)

	// /Dlna
	newDlnaRouter(router)

	// /Environment
	newEnvironmentRouter(router)

	// /Games
	newGamesRouter(router)

	// /GameGenres
	newGameGenresRouter(router)

	// /Genres
	newGenresRouter(router)

	// /Images
	newImagesRouter(router)

	// /Items
	newItemsRouter(router)

	// /Library
	newLibraryRouter(router)

	// /LiveStreams
	newLiveStreamsRouter(router)

	// /LiveTv
	newLiveTvRouter(router)

	// /Localization
	newLocalizationRouter(router)

	// Movies
	newMoviesRouter(router)

	// /MusicGenres
	newMusicGenresRouter(router)

	// /News
	newNewsRouter(router)

	// /Notifications
	newNotificationsRouter(router)

	// /Packages
	newPackagesRouter(router)

	// /Persons
	newPersonsRouter(router)

	// /Playback
	newPlaybackRouter(router)

	// /Playlists
	newPlaylistRouter(router)

	// /Providers
	newProvidersRouter(router)

	// /ScheduledTasks
	newScheduledTasksRouter(router)

	// /Search
	newSearchRouter(router)

	// /Sessions
	newSessionsRouter(router)

	// /Shows
	newShowsRouter(router)

	// /Social
	newSocialRouter(router)

	// /Songs
	newSongsRouter(router)

	// /Startup
	newStartupRouter(router)

	// /Studios
	newStudiosRouter(router)

	// /Sync
	newSyncRouter(router)

	// /System
	newSystemRouter(router)

	// /Trailers
	newTrailersRouter(router)

	// /Users
	newUsersRouter(userManager, router)

	// /Videos
	newVideosRouter(router)

	// /Years
	newYearsRouter(router)

	return router
}
