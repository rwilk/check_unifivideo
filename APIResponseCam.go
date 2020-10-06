package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// APIResponseCam - JSON response with cameras data from UniFi-Video Controller
type APIResponseCam struct {
	Data []CamEntity `json:"data"`
	Meta struct {
		TotalCount    int `json:"totalCount"`
		FilteredCount int `json:"filteredCount"`
	} `json:"meta"`
}

// CamEntity represents single camera in API Response
type CamEntity struct {
	// extra field for link ScheduleEntity API Call
	Schedule *ScheduleEntity

	Name            string `json:"name"`
	UUID            string `json:"uuid"`
	Host            string `json:"host"`
	Model           string `json:"model"`
	Uptime          int64  `json:"uptime"`
	FirmwareVersion string `json:"firmwareVersion"`
	FirmwareBuild   string `json:"firmwareBuild"`
	ProtocolVersion int    `json:"protocolVersion"`
	SystemInfo      struct {
		CPUName string  `json:"cpuName"`
		CPULoad float64 `json:"cpuLoad"`
		Memory  struct {
			Used  int `json:"used"`
			Total int `json:"total"`
		} `json:"memory"`
		AppMemory interface{} `json:"appMemory"`
		Nics      []struct {
			Desc  string `json:"desc"`
			Mac   string `json:"mac"`
			IP    string `json:"ip"`
			RxBps int    `json:"rxBps"`
			TxBps int    `json:"txBps"`
		} `json:"nics"`
		Disk interface{} `json:"disk"`
	} `json:"systemInfo"`
	Mac                      string      `json:"mac"`
	Managed                  bool        `json:"managed"`
	ManagedByOthers          bool        `json:"managedByOthers"`
	Provisioned              bool        `json:"provisioned"`
	UnmanagementRequested    bool        `json:"unmanagementRequested"`
	LastSeen                 int64       `json:"lastSeen"`
	InternalHost             string      `json:"internalHost"`
	State                    string      `json:"state"`
	DisconnectReason         interface{} `json:"disconnectReason"`
	Platform                 string      `json:"platform"`
	ManagementToken          string      `json:"managementToken"`
	ControllerHostAddress    string      `json:"controllerHostAddress"`
	ControllerHostPort       int         `json:"controllerHostPort"`
	Username                 string      `json:"username"`
	LastRecordingID          string      `json:"lastRecordingId"`
	LastRecordingStartTime   int64       `json:"lastRecordingStartTime"`
	RecordingIndicator       string      `json:"recordingIndicator"`
	EnableRecordingIndicator bool        `json:"enableRecordingIndicator"`
	DeviceSettings           struct {
		Name     string      `json:"name"`
		Timezone string      `json:"timezone"`
		Region   interface{} `json:"region"`
		Persists bool        `json:"persists"`
	} `json:"deviceSettings"`
	EnableSuggestedVideoSettings bool `json:"enableSuggestedVideoSettings"`
	MicVolume                    int  `json:"micVolume"`
	AudioBitRate                 int  `json:"audioBitRate"`
	Channels                     []struct {
		ID                       string   `json:"id"`
		Name                     string   `json:"name"`
		Enabled                  bool     `json:"enabled"`
		IsRtspEnabled            bool     `json:"isRtspEnabled"`
		RtspUris                 []string `json:"rtspUris"`
		IsRtmpEnabled            bool     `json:"isRtmpEnabled"`
		RtmpUris                 []string `json:"rtmpUris"`
		IsRtmpsEnabled           bool     `json:"isRtmpsEnabled"`
		RtmpsUris                []string `json:"rtmpsUris"`
		Width                    int      `json:"width"`
		Height                   int      `json:"height"`
		Fps                      int      `json:"fps"`
		Bitrate                  int      `json:"bitrate"`
		MinBitrate               int      `json:"minBitrate"`
		MaxBitrate               int      `json:"maxBitrate"`
		FpsValues                []int    `json:"fpsValues"`
		IdrInterval              int      `json:"idrInterval"`
		IsAdaptiveBitrateEnabled bool     `json:"isAdaptiveBitrateEnabled"`
	} `json:"channels"`
	DefaultStreamingChannel string `json:"defaultStreamingChannel"`
	IspSettings             struct {
		Brightness               int         `json:"brightness"`
		Contrast                 int         `json:"contrast"`
		Denoise                  int         `json:"denoise"`
		Hue                      int         `json:"hue"`
		Saturation               int         `json:"saturation"`
		Sharpness                int         `json:"sharpness"`
		Flip                     int         `json:"flip"`
		Mirror                   int         `json:"mirror"`
		AutoFlipMirror           int         `json:"autoFlipMirror"`
		Gamma                    interface{} `json:"gamma"`
		Wdr                      int         `json:"wdr"`
		AeMode                   string      `json:"aeMode"`
		IrLedMode                string      `json:"irLedMode"`
		IrLedLevel               int         `json:"irLedLevel"`
		FocusMode                string      `json:"focusMode"`
		FocusPosition            int         `json:"focusPosition"`
		ZoomPosition             int         `json:"zoomPosition"`
		IcrSensitivity           int         `json:"icrSensitivity"`
		AggressiveAntiFlicker    int         `json:"aggressiveAntiFlicker"`
		Enable3Dnr               int         `json:"enable3dnr"`
		DZoomStreamID            interface{} `json:"dZoomStreamId"`
		DZoomCenterX             interface{} `json:"dZoomCenterX"`
		DZoomCenterY             interface{} `json:"dZoomCenterY"`
		DZoomScale               interface{} `json:"dZoomScale"`
		LensDistortionCorrection interface{} `json:"lensDistortionCorrection"`
		EnableExternalIr         interface{} `json:"enableExternalIr"`
		TouchFocusX              int         `json:"touchFocusX"`
		TouchFocusY              int         `json:"touchFocusY"`
		IrOnValBrightness        int         `json:"irOnValBrightness"`
		IrOnStsBrightness        int         `json:"irOnStsBrightness"`
		IrOnValContrast          int         `json:"irOnValContrast"`
		IrOnStsContrast          int         `json:"irOnStsContrast"`
		IrOnValDenoise           int         `json:"irOnValDenoise"`
		IrOnStsDenoise           int         `json:"irOnStsDenoise"`
		IrOnValHue               int         `json:"irOnValHue"`
		IrOnStsHue               int         `json:"irOnStsHue"`
		IrOnValSaturation        int         `json:"irOnValSaturation"`
		IrOnStsSaturation        int         `json:"irOnStsSaturation"`
		IrOnValSharpness         int         `json:"irOnValSharpness"`
		IrOnStsSharpness         int         `json:"irOnStsSharpness"`
	} `json:"ispSettings"`
	OsdSettings struct {
		Tag                      string `json:"tag"`
		OverrideMessage          bool   `json:"overrideMessage"`
		EnableDate               int    `json:"enableDate"`
		EnableLogo               int    `json:"enableLogo"`
		EnableStreamerStatsLevel int    `json:"enableStreamerStatsLevel"`
	} `json:"osdSettings"`
	RecordingSettings struct {
		MotionRecordEnabled   bool        `json:"motionRecordEnabled"`
		FullTimeRecordEnabled bool        `json:"fullTimeRecordEnabled"`
		Channel               string      `json:"channel"`
		PrePaddingSecs        int         `json:"prePaddingSecs"`
		PostPaddingSecs       int         `json:"postPaddingSecs"`
		StoragePath           interface{} `json:"storagePath"`
	} `json:"recordingSettings"`
	ScheduleID string `json:"scheduleId"`
	Zones      []struct {
		Name        string      `json:"name"`
		Sensitivity int         `json:"sensitivity"`
		Bitmap      interface{} `json:"bitmap"`
		Coordinates []struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"coordinates"`
		ID string `json:"_id"`
	} `json:"zones"`
	EnablePrivacyMasks bool          `json:"enablePrivacyMasks"`
	PrivacyMasks       []interface{} `json:"privacyMasks"`
	MapSettings        struct {
		X        float64 `json:"x"`
		Y        float64 `json:"y"`
		MapID    string  `json:"mapId"`
		Angle    float64 `json:"angle"`
		Radius   float64 `json:"radius"`
		Rotation float64 `json:"rotation"`
	} `json:"mapSettings"`
	NetworkStatus struct {
		ConnectionState            int         `json:"connectionState"`
		ConnectionStateDescription string      `json:"connectionStateDescription"`
		Essid                      interface{} `json:"essid"`
		Frequency                  int         `json:"frequency"`
		Quality                    int         `json:"quality"`
		QualityMax                 int         `json:"qualityMax"`
		SignalLevel                int         `json:"signalLevel"`
		LinkSpeedMbps              int         `json:"linkSpeedMbps"`
		IPAddress                  string      `json:"ipAddress"`
	} `json:"networkStatus"`
	Status struct {
		RecordingStatus struct {
			Num0 struct {
				MotionRecordingEnabled   bool `json:"motionRecordingEnabled"`
				FullTimeRecordingEnabled bool `json:"fullTimeRecordingEnabled"`
			} `json:"0"`
			Num1 struct {
				MotionRecordingEnabled   bool `json:"motionRecordingEnabled"`
				FullTimeRecordingEnabled bool `json:"fullTimeRecordingEnabled"`
			} `json:"1"`
			Num2 struct {
				MotionRecordingEnabled   bool `json:"motionRecordingEnabled"`
				FullTimeRecordingEnabled bool `json:"fullTimeRecordingEnabled"`
			} `json:"2"`
		} `json:"recordingStatus"`
		ScheduledAction string `json:"scheduledAction"`
		RemoteHost      string `json:"remoteHost"`
		RemotePort      int    `json:"remotePort"`
	} `json:"status"`
	AuthToken struct {
		AuthToken string `json:"authToken"`
	} `json:"authToken"`
	CertSignature         string `json:"certSignature"`
	HasDefaultCredentials bool   `json:"hasDefaultCredentials"`
	AnalyticsSettings     struct {
		EnableSoundAlert   bool   `json:"enableSoundAlert"`
		AnimateLedOnMotion bool   `json:"animateLedOnMotion"`
		SoundAlertVolume   int    `json:"soundAlertVolume"`
		MinimumMotionSecs  int    `json:"minimumMotionSecs"`
		EndMotionAfterSecs int    `json:"endMotionAfterSecs"`
		ConsecutiveRatio   int    `json:"consecutiveRatio"`
		UseNewMotionAlgo   bool   `json:"useNewMotionAlgo"`
		BgModel            string `json:"bgModel"`
	} `json:"analyticsSettings"`
	EnableStatusLed            bool   `json:"enableStatusLed"`
	LedFaceAlwaysOnWhenManaged bool   `json:"ledFaceAlwaysOnWhenManaged"`
	EnableSpeaker              bool   `json:"enableSpeaker"`
	SystemSoundsEnabled        bool   `json:"systemSoundsEnabled"`
	EnableStats                bool   `json:"enableStats"`
	SpeakerVolume              int    `json:"speakerVolume"`
	Deleted                    bool   `json:"deleted"`
	AuthStatus                 string `json:"authStatus"`
	ID                         string `json:"_id"`
}

// ShouldRecord - return true if camera should records
func (cam CamEntity) ShouldRecord() bool {
	// no schedule - should record full time
	if cam.Schedule == nil {
		return true
	}

	nowWeekday := int(time.Now().Weekday()) + 1 // unifi starts count from 1, golang from 0
	for _, days := range cam.Schedule.DaySchedules {
		if days.DayOfWeek == nowWeekday {
			for _, s := range days.ScheduleItems {
				n := time.Now()
				start := time.Date(n.Year(), n.Month(), n.Day(), s.StartHour, s.StartMinute, 0, 0, time.Local)
				end := time.Date(n.Year(), n.Month(), n.Day(), s.EndHour, s.EndMinute, 0, 0, time.Local)
				if n.After(start) && n.Before(end) {
					if s.Action == "FULL_TIME" {
						return true
					} else {
						return false
					}
				}
			}
		}

	}

	// better to warning if something goes wrong
	return true
}

// APIResource - return resource name in API path
func (APIResponseCam) APIResource() string {
	return "camera"
}

// Get - make API call and fill Response
func (a *APIResponseCam) Get() error {
	URL := getURL(a)

	resp, err := http.Get(URL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, a); err != nil {
		return err
	}

	// extra stuff for link schedule API call to this object
	scheduleresponse := &APIResponseSch{}
	if err := scheduleresponse.Get(); err != nil {
		return fmt.Errorf("APIResponseSch.Get():", err)
	}

	var newData []CamEntity

	for _, d := range a.Data {
		if d.ScheduleID != "" {
			d.Schedule = scheduleresponse.GetEntity(d.ScheduleID)
		}
		newData = append(newData, d)
	}

	a.Data = newData

	return nil

}
