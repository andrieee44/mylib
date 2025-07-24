//go:build linux

package input

const (
	// INPUT_PROP_POINTER indicates that the device needs a pointer.
	INPUT_PROP_POINTER = 0x00

	// INPUT_PROP_DIRECT indicates that the device is a direct input device.
	INPUT_PROP_DIRECT = 0x01

	// INPUT_PROP_BUTTONPAD indicates that the device has buttons under
	// the pad.
	INPUT_PROP_BUTTONPAD = 0x02

	// INPUT_PROP_SEMI_MT indicates that the device reports a single
	// rectangular touch area.
	INPUT_PROP_SEMI_MT = 0x03

	// INPUT_PROP_TOPBUTTONPAD indicates that the device places soft buttons
	// at the top of the pad.
	INPUT_PROP_TOPBUTTONPAD = 0x04

	// INPUT_PROP_POINTING_STICK indicates that the device is a pointing
	// stick.
	INPUT_PROP_POINTING_STICK = 0x05

	// INPUT_PROP_ACCELEROMETER indicates that the device has an
	// accelerometer.
	INPUT_PROP_ACCELEROMETER = 0x06

	// INPUT_PROP_MAX is the highest input property code.
	INPUT_PROP_MAX = 0x1f

	// INPUT_PROP_CNT is the count of input properties.
	INPUT_PROP_CNT = INPUT_PROP_MAX + 1

	// EV_SYN is the event type for synchronization.
	// Synchronization events signal boundaries between batches of events.
	EV_SYN = 0x00

	// EV_KEY is the event type for keys and buttons.
	// Reports key/button press and release.
	EV_KEY = 0x01

	// EV_REL is the event type for relative axis movements.
	// Reports changes in position, e.g., mouse or wheel motion.
	EV_REL = 0x02

	// EV_ABS is the event type for absolute axis values.
	// Reports exact device coordinates, e.g., touchscreens or joysticks.
	EV_ABS = 0x03

	// EV_MSC is the event type for miscellaneous events.
	// Device-specific information not covered by other types.
	EV_MSC = 0x04

	// EV_SW is the event type for switch events.
	// Reports binary state changes, e.g., lid open/close.
	EV_SW = 0x05

	// EV_LED is the event type for LED indicators.
	// Controls device LEDs, e.g., keyboard caps-lock light.
	EV_LED = 0x11

	// EV_SND is the event type for sound events.
	// Triggers device-embedded sound generators.
	EV_SND = 0x12

	// EV_REP is the event type for auto-repeat settings.
	// Adjusts keyboard repeat rates and delays.
	EV_REP = 0x14

	// EV_FF is the event type for force-feedback effects.
	// Starts, stops, or updates rumble/vibration motors.
	EV_FF = 0x15

	// EV_PWR is the event type for power management.
	// Reports power state changes, e.g., battery warnings.
	EV_PWR = 0x16

	// EV_FF_STATUS is the event type for force-feedback status.
	// Reports effect completion or playback errors.
	EV_FF_STATUS = 0x17

	// EV_MAX is the highest defined event type code.
	EV_MAX = 0x1f

	// EV_CNT is the total number of event types.
	EV_CNT = EV_MAX + 1

	// SYN_REPORT marks the end of a batch of input events.
	SYN_REPORT = 0

	// SYN_CONFIG signals a change in device configuration.
	SYN_CONFIG = 1

	// SYN_MT_REPORT reports an update to the multi-touch state.
	SYN_MT_REPORT = 2

	// SYN_DROPPED indicates that one or more input events were dropped.
	SYN_DROPPED = 3

	// SYN_MAX is the highest synchronization event code.
	SYN_MAX = 0x0f

	// SYN_CNT is the total number of synchronization events.
	SYN_CNT = SYN_MAX + 1

	// KEY_RESERVED is reserved (no key).
	KEY_RESERVED = 0

	// KEY_ESC is the Escape key.
	KEY_ESC = 1

	// KEY_1 is the '1' key.
	KEY_1 = 2

	// KEY_2 is the '2' key.
	KEY_2 = 3

	// KEY_3 is the '3' key.
	KEY_3 = 4

	// KEY_4 is the '4' key.
	KEY_4 = 5

	// KEY_5 is the '5' key.
	KEY_5 = 6

	// KEY_6 is the '6' key.
	KEY_6 = 7

	// KEY_7 is the '7' key.
	KEY_7 = 8

	// KEY_8 is the '8' key.
	KEY_8 = 9

	// KEY_9 is the '9' key.
	KEY_9 = 10

	// KEY_0 is the '0' key.
	KEY_0 = 11

	// KEY_MINUS is the '-' key.
	KEY_MINUS = 12

	// KEY_EQUAL is the '=' key.
	KEY_EQUAL = 13

	// KEY_BACKSPACE is the Backspace key.
	KEY_BACKSPACE = 14

	// KEY_TAB is the Tab key.
	KEY_TAB = 15

	// KEY_Q is the 'Q' key.
	KEY_Q = 16

	// KEY_W is the 'W' key.
	KEY_W = 17

	// KEY_E is the 'E' key.
	KEY_E = 18

	// KEY_R is the 'R' key.
	KEY_R = 19

	// KEY_T is the 'T' key.
	KEY_T = 20

	// KEY_Y is the 'Y' key.
	KEY_Y = 21

	// KEY_U is the 'U' key.
	KEY_U = 22

	// KEY_I is the 'I' key.
	KEY_I = 23

	// KEY_O is the 'O' key.
	KEY_O = 24

	// KEY_P is the 'P' key.
	KEY_P = 25

	// KEY_LEFTBRACE is the '[' key.
	KEY_LEFTBRACE = 26

	// KEY_RIGHTBRACE is the ']' key.
	KEY_RIGHTBRACE = 27

	// KEY_ENTER is the Enter key.
	KEY_ENTER = 28

	// KEY_LEFTCTRL is the left Control key.
	KEY_LEFTCTRL = 29

	// KEY_A is the 'A' key.
	KEY_A = 30

	// KEY_S is the 'S' key.
	KEY_S = 31

	// KEY_D is the 'D' key.
	KEY_D = 32

	// KEY_F is the 'F' key.
	KEY_F = 33

	// KEY_G is the 'G' key.
	KEY_G = 34

	// KEY_H is the 'H' key.
	KEY_H = 35

	// KEY_J is the 'J' key.
	KEY_J = 36

	// KEY_K is the 'K' key.
	KEY_K = 37

	// KEY_L is the 'L' key.
	KEY_L = 38

	// KEY_SEMICOLON is the ';' key.
	KEY_SEMICOLON = 39

	// KEY_APOSTROPHE is the ”' key.
	KEY_APOSTROPHE = 40

	// KEY_GRAVE is the '`' key.
	KEY_GRAVE = 41

	// KEY_LEFTSHIFT is the left Shift key.
	KEY_LEFTSHIFT = 42

	// KEY_BACKSLASH is the '\' key.
	KEY_BACKSLASH = 43

	// KEY_Z is the 'Z' key.
	KEY_Z = 44

	// KEY_X is the 'X' key.
	KEY_X = 45

	// KEY_C is the 'C' key.
	KEY_C = 46

	// KEY_V is the 'V' key.
	KEY_V = 47

	// KEY_B is the 'B' key.
	KEY_B = 48

	// KEY_N is the 'N' key.
	KEY_N = 49

	// KEY_M is the 'M' key.
	KEY_M = 50

	// KEY_COMMA is the ',' key.
	KEY_COMMA = 51

	// KEY_DOT is the '.' key.
	KEY_DOT = 52

	// KEY_SLASH is the '/' key.
	KEY_SLASH = 53

	// KEY_RIGHTSHIFT is the right Shift key.
	KEY_RIGHTSHIFT = 54

	// KEY_KPASTERISK is the keypad '*' key.
	KEY_KPASTERISK = 55

	// KEY_LEFTALT is the left Alt key.
	KEY_LEFTALT = 56

	// KEY_SPACE is the Space Bar.
	KEY_SPACE = 57

	// KEY_CAPSLOCK is the Caps Lock key.
	KEY_CAPSLOCK = 58

	// KEY_F1 is the F1 function key.
	KEY_F1 = 59

	// KEY_F2 is the F2 function key.
	KEY_F2 = 60

	// KEY_F3 is the F3 function key.
	KEY_F3 = 61

	// KEY_F4 is the F4 function key.
	KEY_F4 = 62

	// KEY_F5 is the F5 function key.
	KEY_F5 = 63

	// KEY_F6 is the F6 function key.
	KEY_F6 = 64

	// KEY_F7 is the F7 function key.
	KEY_F7 = 65

	// KEY_F8 is the F8 function key.
	KEY_F8 = 66

	// KEY_F9 is the F9 function key.
	KEY_F9 = 67

	// KEY_F10 is the F10 function key.
	KEY_F10 = 68

	// KEY_NUMLOCK is the Num Lock key.
	KEY_NUMLOCK = 69

	// KEY_SCROLLLOCK is the Scroll Lock key.
	KEY_SCROLLLOCK = 70

	// KEY_KP7 is the keypad '7' key.
	KEY_KP7 = 71

	// KEY_KP8 is the keypad '8' key.
	KEY_KP8 = 72

	// KEY_KP9 is the keypad '9' key.
	KEY_KP9 = 73

	// KEY_KPMINUS is the keypad '-' key.
	KEY_KPMINUS = 74

	// KEY_KP4 is the keypad '4' key.
	KEY_KP4 = 75

	// KEY_KP5 is the keypad '5' key.
	KEY_KP5 = 76

	// KEY_KP6 is the keypad '6' key.
	KEY_KP6 = 77

	// KEY_KPPLUS is the keypad '+' key.
	KEY_KPPLUS = 78

	// KEY_KP1 is the keypad '1' key.
	KEY_KP1 = 79

	// KEY_KP2 is the keypad '2' key.
	KEY_KP2 = 80

	// KEY_KP3 is the keypad '3' key.
	KEY_KP3 = 81

	// KEY_KP0 is the keypad '0' key.
	KEY_KP0 = 82

	// KEY_KPDOT is the keypad '.' key.
	KEY_KPDOT = 83

	// KEY_ZENKAKUHANKAKU is the Zenkaku/Hankaku key.
	KEY_ZENKAKUHANKAKU = 85

	// KEY_102ND is the ISO 102nd key.
	KEY_102ND = 86

	// KEY_F11 is the F11 function key.
	KEY_F11 = 87

	// KEY_F12 is the F12 function key.
	KEY_F12 = 88

	// KEY_RO is the RO key (Japanese keyboard).
	KEY_RO = 89

	// KEY_KATAKANA is the Katakana key.
	KEY_KATAKANA = 90

	// KEY_HIRAGANA is the Hiragana key.
	KEY_HIRAGANA = 91

	// KEY_HENKAN is the Henkan (conversion) key.
	KEY_HENKAN = 92

	// KEY_KATAKANAHIRAGANA toggles between Katakana and Hiragana.
	KEY_KATAKANAHIRAGANA = 93

	// KEY_MUHENKAN is the Muhenkan (non-conversion) key.
	KEY_MUHENKAN = 94

	// KEY_KPJPCOMMA is the keypad Japanese comma key.
	KEY_KPJPCOMMA = 95

	// KEY_KPENTER is the keypad Enter key.
	KEY_KPENTER = 96

	// KEY_RIGHTCTRL is the right Control key.
	KEY_RIGHTCTRL = 97

	// KEY_KPSLASH is the keypad slash key.
	KEY_KPSLASH = 98

	// KEY_SYSRQ is the System Request (SysRq) key.
	KEY_SYSRQ = 99

	// KEY_RIGHTALT is the right Alt key.
	KEY_RIGHTALT = 100

	// KEY_LINEFEED is the Line Feed key.
	KEY_LINEFEED = 101

	// KEY_HOME is the Home key.
	KEY_HOME = 102

	// KEY_UP is the Up Arrow key.
	KEY_UP = 103

	// KEY_PAGEUP is the Page Up key.
	KEY_PAGEUP = 104

	// KEY_LEFT is the Left Arrow key.
	KEY_LEFT = 105

	// KEY_RIGHT is the Right Arrow key.
	KEY_RIGHT = 106

	// KEY_END is the End key.
	KEY_END = 107

	// KEY_DOWN is the Down Arrow key.
	KEY_DOWN = 108

	// KEY_PAGEDOWN is the Page Down key.
	KEY_PAGEDOWN = 109

	// KEY_INSERT is the Insert key.
	KEY_INSERT = 110

	// KEY_DELETE is the Delete key.
	KEY_DELETE = 111

	// KEY_MACRO is the Macro key.
	KEY_MACRO = 112

	// KEY_MUTE is the Mute key.
	KEY_MUTE = 113

	// KEY_VOLUMEDOWN is the Volume Down key.
	KEY_VOLUMEDOWN = 114

	// KEY_VOLUMEUP is the Volume Up key.
	KEY_VOLUMEUP = 115

	// KEY_POWER is the System Power Down key.
	KEY_POWER = 116

	// KEY_KPEQUAL is the keypad Equal key.
	KEY_KPEQUAL = 117

	// KEY_KPPLUSMINUS is the keypad Plus/Minus key.
	KEY_KPPLUSMINUS = 118

	// KEY_PAUSE is the Pause key.
	KEY_PAUSE = 119

	// KEY_SCALE is the Compiz Scale (Expose) key.
	KEY_SCALE = 120

	// KEY_KPCOMMA is the keypad comma key.
	KEY_KPCOMMA = 121

	// KEY_HANGEUL is the Hangeul key (Korean).
	KEY_HANGEUL = 122

	// KEY_HANGUEL is an alias for KEY_HANGEUL.
	KEY_HANGUEL = KEY_HANGEUL

	// KEY_HANJA is the Hanja key (Korean).
	KEY_HANJA = 123

	// KEY_YEN is the Yen currency key.
	KEY_YEN = 124

	// KEY_LEFTMETA is the left Meta (Windows) key.
	KEY_LEFTMETA = 125

	// KEY_RIGHTMETA is the right Meta (Windows) key.
	KEY_RIGHTMETA = 126

	// KEY_COMPOSE is the Compose key.
	KEY_COMPOSE = 127

	// KEY_STOP is the Application Control Stop key.
	KEY_STOP = 128

	// KEY_AGAIN is the Application Control Again key.
	KEY_AGAIN = 129

	// KEY_PROPS is the Application Control Properties key.
	KEY_PROPS = 130

	// KEY_UNDO is the Application Control Undo key.
	KEY_UNDO = 131

	// KEY_FRONT is the Application Control Front key.
	KEY_FRONT = 132

	// KEY_COPY is the Application Control Copy key.
	KEY_COPY = 133

	// KEY_OPEN is the Application Control Open key.
	KEY_OPEN = 134

	// KEY_PASTE is the Application Control Paste key.
	KEY_PASTE = 135

	// KEY_FIND is the Application Control Search key.
	KEY_FIND = 136

	// KEY_CUT is the Application Control Cut key.
	KEY_CUT = 137

	// KEY_HELP is the Application Launch Integrated Help Center key.
	KEY_HELP = 138

	// KEY_MENU is the Application Launch Menu key.
	KEY_MENU = 139

	// KEY_CALC is the Application Launch Calculator key.
	KEY_CALC = 140

	// KEY_SETUP is the Application Control Setup key.
	KEY_SETUP = 141

	// KEY_SLEEP is the System Control Sleep key.
	KEY_SLEEP = 142

	// KEY_WAKEUP is the System Control Wake Up key.
	KEY_WAKEUP = 143

	// KEY_FILE is the Application Launch Local Machine Browser key.
	KEY_FILE = 144

	// KEY_SENDFILE is the Application Control Send File key.
	KEY_SENDFILE = 145

	// KEY_DELETEFILE is the Application Control Delete File key.
	KEY_DELETEFILE = 146

	// KEY_XFER is the Application Control Transfer key.
	KEY_XFER = 147

	// KEY_PROG1 is the Application Launch Program 1 key.
	KEY_PROG1 = 148

	// KEY_PROG2 is the Application Launch Program 2 key.
	KEY_PROG2 = 149

	// KEY_WWW is the Application Launch Internet Browser key.
	KEY_WWW = 150

	// KEY_MSDOS is the Application Launch MSDOS key.
	KEY_MSDOS = 151

	// KEY_COFFEE is the Application Launch Terminal Lock/Screensaver key.
	KEY_COFFEE = 152

	// KEY_SCREENLOCK is an alias for KEY_COFFEE.
	KEY_SCREENLOCK = KEY_COFFEE

	// KEY_ROTATE_DISPLAY is the Application Control Display Orientation
	// key.
	KEY_ROTATE_DISPLAY = 153

	// KEY_DIRECTION is an alias for KEY_ROTATE_DISPLAY.
	KEY_DIRECTION = KEY_ROTATE_DISPLAY

	// KEY_CYCLEWINDOWS is the Application Control Cycle Windows key.
	KEY_CYCLEWINDOWS = 154

	// KEY_MAIL is the Application Launch Mail key.
	KEY_MAIL = 155

	// KEY_BOOKMARKS is the Application Control Bookmarks key.
	KEY_BOOKMARKS = 156

	// KEY_COMPUTER is the Application Launch Computer key.
	KEY_COMPUTER = 157

	// KEY_BACK is the Application Control Back key.
	KEY_BACK = 158

	// KEY_FORWARD is the Application Control Forward key.
	KEY_FORWARD = 159

	// KEY_CLOSECD is the Application Control Close CD key.
	KEY_CLOSECD = 160

	// KEY_EJECTCD is the Application Control Eject CD key.
	KEY_EJECTCD = 161

	// KEY_EJECTCLOSECD is the Application Control Eject/Close CD key.
	KEY_EJECTCLOSECD = 162

	// KEY_NEXTSONG is the Application Control Next Song key.
	KEY_NEXTSONG = 163

	// KEY_PLAYPAUSE is the Application Control Play/Pause key.
	KEY_PLAYPAUSE = 164

	// KEY_PREVIOUSSONG is the Application Control Previous Song key.
	KEY_PREVIOUSSONG = 165

	// KEY_STOPCD is the Application Control Stop CD key.
	KEY_STOPCD = 166

	// KEY_RECORD is the Application Control Record key.
	KEY_RECORD = 167

	// KEY_REWIND is the Application Control Rewind key.
	KEY_REWIND = 168

	// KEY_PHONE is the Application Launch Telephone key.
	KEY_PHONE = 169

	// KEY_ISO is the ISO special function key.
	KEY_ISO = 170

	// KEY_CONFIG is the Application Launch Consumer Control
	// Configuration key.
	KEY_CONFIG = 171

	// KEY_HOMEPAGE is the Application Control Home key.
	KEY_HOMEPAGE = 172

	// KEY_REFRESH is the Application Control Refresh key.
	KEY_REFRESH = 173

	// KEY_EXIT is the Application Control Exit key.
	KEY_EXIT = 174

	// KEY_MOVE is the Application Control Move key.
	KEY_MOVE = 175

	// KEY_EDIT is the Application Control Edit key.
	KEY_EDIT = 176

	// KEY_SCROLLUP is the Application Control Scroll Up key.
	KEY_SCROLLUP = 177

	// KEY_SCROLLDOWN is the Application Control Scroll Down key.
	KEY_SCROLLDOWN = 178

	// KEY_KPLEFTPAREN is the keypad '(' key.
	KEY_KPLEFTPAREN = 179

	// KEY_KPRIGHTPAREN is the keypad ')' key.
	KEY_KPRIGHTPAREN = 180

	// KEY_NEW is the Application Control New key.
	KEY_NEW = 181

	// KEY_REDO is the Application Control Redo/Repeat key.
	KEY_REDO = 182

	// KEY_F13 is the F13 function key.
	KEY_F13 = 183

	// KEY_F14 is the F14 function key.
	KEY_F14 = 184

	// KEY_F15 is the F15 function key.
	KEY_F15 = 185

	// KEY_F16 is the F16 function key.
	KEY_F16 = 186

	// KEY_F17 is the F17 function key.
	KEY_F17 = 187

	// KEY_F18 is the F18 function key.
	KEY_F18 = 188

	// KEY_F19 is the F19 function key.
	KEY_F19 = 189

	// KEY_F20 is the F20 function key.
	KEY_F20 = 190

	// KEY_F21 is the F21 function key.
	KEY_F21 = 191

	// KEY_F22 is the F22 function key.
	KEY_F22 = 192

	// KEY_F23 is the F23 function key.
	KEY_F23 = 193

	// KEY_F24 is the F24 function key.
	KEY_F24 = 194

	// KEY_PLAYCD is the Play CD key.
	KEY_PLAYCD = 200

	// KEY_PAUSECD is the Pause CD key.
	KEY_PAUSECD = 201

	// KEY_PROG3 is the Program 3 key.
	KEY_PROG3 = 202

	// KEY_PROG4 is the Program 4 key.
	KEY_PROG4 = 203

	// KEY_ALL_APPLICATIONS is the Application Control Desktop Show
	// All Applications key.
	KEY_ALL_APPLICATIONS = 204

	// KEY_DASHBOARD is an alias for KEY_ALL_APPLICATIONS.
	KEY_DASHBOARD = KEY_ALL_APPLICATIONS

	// KEY_SUSPEND is the System Control Suspend key.
	KEY_SUSPEND = 205

	// KEY_CLOSE is the Application Control Close key.
	KEY_CLOSE = 206

	// KEY_PLAY is the Play key.
	KEY_PLAY = 207

	// KEY_FASTFORWARD is the Fast Forward key.
	KEY_FASTFORWARD = 208

	// KEY_BASSBOOST is the Bass Boost key.
	KEY_BASSBOOST = 209

	// KEY_PRINT is the Application Control Print key.
	KEY_PRINT = 210

	// KEY_HP is the HP key.
	KEY_HP = 211

	// KEY_CAMERA is the Camera key.
	KEY_CAMERA = 212

	// KEY_SOUND is the Sound key.
	KEY_SOUND = 213

	// KEY_QUESTION is the Question key.
	KEY_QUESTION = 214

	// KEY_EMAIL is the Email key.
	KEY_EMAIL = 215

	// KEY_CHAT is the Chat key.
	KEY_CHAT = 216

	// KEY_SEARCH is the Search key.
	KEY_SEARCH = 217

	// KEY_CONNECT is the Connect key.
	KEY_CONNECT = 218

	// KEY_FINANCE is the Application Launch Checkbook/Finance key.
	KEY_FINANCE = 219

	// KEY_SPORT is the Sport key.
	KEY_SPORT = 220

	// KEY_SHOP is the Shop key.
	KEY_SHOP = 221

	// KEY_ALTERASE is the Alternate Erase key.
	KEY_ALTERASE = 222

	// KEY_CANCEL is the Application Control Cancel key.
	KEY_CANCEL = 223

	// KEY_BRIGHTNESSDOWN is the Brightness Down key.
	KEY_BRIGHTNESSDOWN = 224

	// KEY_BRIGHTNESSUP is the Brightness Up key.
	KEY_BRIGHTNESSUP = 225

	// KEY_MEDIA is the Media key.
	KEY_MEDIA = 226

	// KEY_SWITCHVIDEOMODE cycles between available video outputs
	// (monitor, LCD, TV-out, etc).
	KEY_SWITCHVIDEOMODE = 227

	// KEY_KBDILLUMTOGGLE toggles the keyboard illumination on and off.
	KEY_KBDILLUMTOGGLE = 228

	// KEY_KBDILLUMDOWN decreases the keyboard illumination.
	KEY_KBDILLUMDOWN = 229

	// KEY_KBDILLUMUP increases the keyboard illumination.
	KEY_KBDILLUMUP = 230

	// KEY_SEND is the Application Control Send key.
	KEY_SEND = 231

	// KEY_REPLY is the Application Control Reply key.
	KEY_REPLY = 232

	// KEY_FORWARDMAIL is the Application Control Forward Msg key.
	KEY_FORWARDMAIL = 233

	// KEY_SAVE is the Application Control Save key.
	KEY_SAVE = 234

	// KEY_DOCUMENTS is the Documents key.
	KEY_DOCUMENTS = 235

	// KEY_BATTERY is the Battery key.
	KEY_BATTERY = 236

	// KEY_BLUETOOTH is the Bluetooth key.
	KEY_BLUETOOTH = 237

	// KEY_WLAN is the WLAN key.
	KEY_WLAN = 238

	// KEY_UWB is the UWB key.
	KEY_UWB = 239

	// KEY_UNKNOWN is the Unknown key.
	KEY_UNKNOWN = 240

	// KEY_VIDEO_NEXT is the drive next video source key.
	KEY_VIDEO_NEXT = 241

	// KEY_VIDEO_PREV is the drive previous video source key.
	KEY_VIDEO_PREV = 242

	// KEY_BRIGHTNESS_CYCLE is the Brightness Cycle key; brightness
	// increases and wraps to minimum after the maximum.
	KEY_BRIGHTNESS_CYCLE = 243

	// KEY_BRIGHTNESS_AUTO sets auto brightness; manual brightness control
	// is off and relies on ambient light.
	KEY_BRIGHTNESS_AUTO = 244

	// KEY_BRIGHTNESS_ZERO is an alias for KEY_BRIGHTNESS_AUTO.
	KEY_BRIGHTNESS_ZERO = KEY_BRIGHTNESS_AUTO

	// KEY_DISPLAY_OFF sets the display device to off state.
	KEY_DISPLAY_OFF = 245

	// KEY_WWAN is the Wireless WAN key (LTE, UMTS, GSM, etc.).
	KEY_WWAN = 246

	// KEY_WIMAX is an alias for the KEY_WWAN.
	KEY_WIMAX = KEY_WWAN

	// KEY_RFKILL is the key that controls all radios
	// (WiFi, Bluetooth, etc.).
	KEY_RFKILL = 247

	// KEY_MICMUTE toggles microphone mute/unmute.
	KEY_MICMUTE = 248

	// BTN_MISC marks the start of miscellaneous button codes.
	BTN_MISC = 0x100

	// BTN_0 is an alias for BTN_MISC, representing the first
	// miscellaneous button.
	BTN_0 = 0x100

	// BTN_1 is the second miscellaneous button.
	BTN_1 = 0x101

	// BTN_2 is the third miscellaneous button.
	BTN_2 = 0x102

	// BTN_3 is the fourth miscellaneous button.
	BTN_3 = 0x103

	// BTN_4 is the fifth miscellaneous button.
	BTN_4 = 0x104

	// BTN_5 is the sixth miscellaneous button.
	BTN_5 = 0x105

	// BTN_6 is the seventh miscellaneous button.
	BTN_6 = 0x106

	// BTN_7 is the eighth miscellaneous button.
	BTN_7 = 0x107

	// BTN_8 is the ninth miscellaneous button.
	BTN_8 = 0x108

	// BTN_9 is the tenth miscellaneous button.
	BTN_9 = 0x109

	// BTN_MOUSE marks the start of mouse button codes.
	BTN_MOUSE = 0x110

	// BTN_LEFT is the left mouse button.
	BTN_LEFT = 0x110

	// BTN_RIGHT is the right mouse button.
	BTN_RIGHT = 0x111

	// BTN_MIDDLE is the middle mouse button.
	BTN_MIDDLE = 0x112

	// BTN_SIDE is the side mouse button.
	BTN_SIDE = 0x113

	// BTN_EXTRA is an extra mouse button.
	BTN_EXTRA = 0x114

	// BTN_FORWARD is the forward mouse button.
	BTN_FORWARD = 0x115

	// BTN_BACK is the back mouse button.
	BTN_BACK = 0x116

	// BTN_TASK is the task mouse button.
	BTN_TASK = 0x117

	// BTN_JOYSTICK marks the start of joystick button codes.
	BTN_JOYSTICK = 0x120

	// BTN_TRIGGER is the primary trigger button on a joystick.
	BTN_TRIGGER = 0x120

	// BTN_THUMB is the thumb button on a joystick.
	BTN_THUMB = 0x121

	// BTN_THUMB2 is the second thumb button on a joystick.
	BTN_THUMB2 = 0x122

	// BTN_TOP is the top button on a joystick.
	BTN_TOP = 0x123

	// BTN_TOP2 is the second top button on a joystick.
	BTN_TOP2 = 0x124

	// BTN_PINKIE is the pinkie (little finger) button on a joystick.
	BTN_PINKIE = 0x125

	// BTN_BASE is the first base button on a joystick.
	BTN_BASE = 0x126

	// BTN_BASE2 is the second base button on a joystick.
	BTN_BASE2 = 0x127

	// BTN_BASE3 is the third base button on a joystick.
	BTN_BASE3 = 0x128

	// BTN_BASE4 is the fourth base button on a joystick.
	BTN_BASE4 = 0x129

	// BTN_BASE5 is the fifth base button on a joystick.
	BTN_BASE5 = 0x12a

	// BTN_BASE6 is the sixth base button on a joystick.
	BTN_BASE6 = 0x12b

	// BTN_DEAD is reserved (no assigned button).
	BTN_DEAD = 0x12f

	// BTN_GAMEPAD is the first gamepad button code.
	BTN_GAMEPAD = 0x130

	// BTN_SOUTH is the South face button on a gamepad.
	BTN_SOUTH = 0x130

	// BTN_A is an alias for the South face button.
	BTN_A = BTN_SOUTH

	// BTN_EAST is the East face button on a gamepad.
	BTN_EAST = 0x131

	// BTN_B is an alias for the East face button.
	BTN_B = BTN_EAST

	// BTN_C is the C face button on a gamepad.
	BTN_C = 0x132

	// BTN_NORTH is the North face button on a gamepad.
	BTN_NORTH = 0x133

	// BTN_X is an alias for the North face button.
	BTN_X = BTN_NORTH

	// BTN_WEST is the West face button on a gamepad.
	BTN_WEST = 0x134

	// BTN_Y is an alias for the West face button.
	BTN_Y = BTN_WEST

	// BTN_Z is the Z face button on a gamepad.
	BTN_Z = 0x135

	// BTN_TL is the top left shoulder button.
	BTN_TL = 0x136

	// BTN_TR is the top right shoulder button.
	BTN_TR = 0x137

	// BTN_TL2 is the second top left shoulder button.
	BTN_TL2 = 0x138

	// BTN_TR2 is the second top right shoulder button.
	BTN_TR2 = 0x139

	// BTN_SELECT is the Select button.
	BTN_SELECT = 0x13a

	// BTN_START is the Start button.
	BTN_START = 0x13b

	// BTN_MODE is the Mode button.
	BTN_MODE = 0x13c

	// BTN_THUMBL is the left thumb button.
	BTN_THUMBL = 0x13d

	// BTN_THUMBR is the right thumb button.
	BTN_THUMBR = 0x13e

	// BTN_DIGI marks the start of digitizer tool codes.
	BTN_DIGI = 0x140

	// BTN_TOOL_PEN is the pen tool.
	BTN_TOOL_PEN = 0x140

	// BTN_TOOL_RUBBER is the rubber (eraser) tool.
	BTN_TOOL_RUBBER = 0x141

	// BTN_TOOL_BRUSH is the brush tool.
	BTN_TOOL_BRUSH = 0x142

	// BTN_TOOL_PENCIL is the pencil tool.
	BTN_TOOL_PENCIL = 0x143

	// BTN_TOOL_AIRBRUSH is the airbrush tool.
	BTN_TOOL_AIRBRUSH = 0x144

	// BTN_TOOL_FINGER is the finger touch tool.
	BTN_TOOL_FINGER = 0x145

	// BTN_TOOL_MOUSE is the mouse tool.
	BTN_TOOL_MOUSE = 0x146

	// BTN_TOOL_LENS is the lens tool.
	BTN_TOOL_LENS = 0x147

	// BTN_TOOL_QUINTTAP detects five fingers on a trackpad.
	BTN_TOOL_QUINTTAP = 0x148

	// BTN_STYLUS3 is the third stylus button.
	BTN_STYLUS3 = 0x149

	// BTN_TOUCH indicates a touch event on the digitizer.
	BTN_TOUCH = 0x14a

	// BTN_STYLUS is the primary stylus tool.
	BTN_STYLUS = 0x14b

	// BTN_STYLUS2 is the secondary stylus tool.
	BTN_STYLUS2 = 0x14c

	// BTN_TOOL_DOUBLETAP detects two fingers on a trackpad.
	BTN_TOOL_DOUBLETAP = 0x14d

	// BTN_TOOL_TRIPLETAP detects three fingers on a trackpad.
	BTN_TOOL_TRIPLETAP = 0x14e

	// BTN_TOOL_QUADTAP detects four fingers on a trackpad.
	BTN_TOOL_QUADTAP = 0x14f

	// BTN_WHEEL is the wheel button.
	BTN_WHEEL = 0x150

	// BTN_GEAR_DOWN is an alias for BTN_WHEEL.
	BTN_GEAR_DOWN = BTN_WHEEL

	// BTN_GEAR_UP is the gear up button.
	BTN_GEAR_UP = 0x151

	// KEY_OK is the OK key.
	KEY_OK = 0x160

	// KEY_SELECT is the Select key.
	KEY_SELECT = 0x161

	// KEY_GOTO is the Goto key.
	KEY_GOTO = 0x162

	// KEY_CLEAR is the Clear key.
	KEY_CLEAR = 0x163

	// KEY_POWER2 is the second Power key.
	KEY_POWER2 = 0x164

	// KEY_OPTION is the Option key.
	KEY_OPTION = 0x165

	// KEY_INFO is the Application Launch OEM Features/Tips/Tutorial key.
	KEY_INFO = 0x166

	// KEY_TIME is the Time key.
	KEY_TIME = 0x167

	// KEY_VENDOR is the Vendor key.
	KEY_VENDOR = 0x168

	// KEY_ARCHIVE is the Archive key.
	KEY_ARCHIVE = 0x169

	// KEY_PROGRAM is the Media Select Program Guide key.
	KEY_PROGRAM = 0x16a

	// KEY_CHANNEL is the Channel key.
	KEY_CHANNEL = 0x16b

	// KEY_FAVORITES is the Favorites key.
	KEY_FAVORITES = 0x16c

	// KEY_EPG is the EPG key.
	KEY_EPG = 0x16d

	// KEY_PVR is the Media Select Home key.
	KEY_PVR = 0x16e

	// KEY_MHP is the MHP key.
	KEY_MHP = 0x16f

	// KEY_LANGUAGE is the Language key.
	KEY_LANGUAGE = 0x170

	// KEY_TITLE is the Title key.
	KEY_TITLE = 0x171

	// KEY_SUBTITLE is the Subtitle key.
	KEY_SUBTITLE = 0x172

	// KEY_ANGLE is the Angle key.
	KEY_ANGLE = 0x173

	// KEY_FULL_SCREEN is the Application Control View Toggle key.
	KEY_FULL_SCREEN = 0x174

	// KEY_ZOOM is an alias for KEY_FULL_SCREEN.
	KEY_ZOOM = KEY_FULL_SCREEN

	// KEY_MODE is the Mode key.
	KEY_MODE = 0x175

	// KEY_KEYBOARD is the Keyboard key.
	KEY_KEYBOARD = 0x176

	// KEY_ASPECT_RATIO is the Aspect Ratio key.
	KEY_ASPECT_RATIO = 0x177

	// KEY_SCREEN is an alias for KEY_ASPECT_RATIO.
	KEY_SCREEN = KEY_ASPECT_RATIO

	// KEY_PC is the Media Select Computer key.
	KEY_PC = 0x178

	// KEY_TV is the Media Select TV key.
	KEY_TV = 0x179

	// KEY_TV2 is the Media Select Cable key.
	KEY_TV2 = 0x17a

	// KEY_VCR is the Media Select VCR key.
	KEY_VCR = 0x17b

	// KEY_VCR2 is the VCR Plus key.
	KEY_VCR2 = 0x17c

	// KEY_SAT is the Media Select Satellite key.
	KEY_SAT = 0x17d

	// KEY_SAT2 is the second Satellite key.
	KEY_SAT2 = 0x17e

	// KEY_CD is the Media Select CD key.
	KEY_CD = 0x17f

	// KEY_TAPE is the Media Select Tape key.
	KEY_TAPE = 0x180

	// KEY_RADIO is the Radio key.
	KEY_RADIO = 0x181

	// KEY_TUNER is the Media Select Tuner key.
	KEY_TUNER = 0x182

	// KEY_PLAYER is the Player key.
	KEY_PLAYER = 0x183

	// KEY_TEXT is the Text key.
	KEY_TEXT = 0x184

	// KEY_DVD is the Media Select DVD key.
	KEY_DVD = 0x185

	// KEY_AUX is the Aux key.
	KEY_AUX = 0x186

	// KEY_MP3 is the MP3 key.
	KEY_MP3 = 0x187

	// KEY_AUDIO is the Application Launch Audio Browser key.
	KEY_AUDIO = 0x188

	// KEY_VIDEO is the Application Launch Movie Browser key.
	KEY_VIDEO = 0x189

	// KEY_DIRECTORY is the Directory key.
	KEY_DIRECTORY = 0x18a

	// KEY_LIST is the List key.
	KEY_LIST = 0x18b

	// KEY_MEMO is the Media Select Messages key.
	KEY_MEMO = 0x18c

	// KEY_CALENDAR is the Calendar key.
	KEY_CALENDAR = 0x18d

	// KEY_RED is the Red key.
	KEY_RED = 0x18e

	// KEY_GREEN is the Green key.
	KEY_GREEN = 0x18f

	// KEY_YELLOW is the Yellow key.
	KEY_YELLOW = 0x190

	// KEY_BLUE is the Blue key.
	KEY_BLUE = 0x191

	// KEY_CHANNELUP is the Channel Increment key.
	KEY_CHANNELUP = 0x192

	// KEY_CHANNELDOWN is the Channel Decrement key.
	KEY_CHANNELDOWN = 0x193

	// KEY_FIRST is the First key.
	KEY_FIRST = 0x194

	// KEY_LAST is the Recall Last key.
	KEY_LAST = 0x195

	// KEY_AB is the AB key.
	KEY_AB = 0x196

	// KEY_NEXT is the Next key.
	KEY_NEXT = 0x197

	// KEY_RESTART is the Restart key.
	KEY_RESTART = 0x198

	// KEY_SLOW is the Slow key.
	KEY_SLOW = 0x199

	// KEY_SHUFFLE is the Shuffle key.
	KEY_SHUFFLE = 0x19a

	// KEY_BREAK is the Break key.
	KEY_BREAK = 0x19b

	// KEY_PREVIOUS is the Previous key.
	KEY_PREVIOUS = 0x19c

	// KEY_DIGITS is the Digits key.
	KEY_DIGITS = 0x19d

	// KEY_TEEN is the Teen key.
	KEY_TEEN = 0x19e

	// KEY_TWEN is the Twen key.
	KEY_TWEN = 0x19f

	// KEY_VIDEOPHONE is the Media Select Video Phone key.
	KEY_VIDEOPHONE = 0x1a0

	// KEY_GAMES is the Media Select Games key.
	KEY_GAMES = 0x1a1

	// KEY_ZOOMIN is the Application Control Zoom In key.
	KEY_ZOOMIN = 0x1a2

	// KEY_ZOOMOUT is the Application Control Zoom Out key.
	KEY_ZOOMOUT = 0x1a3

	// KEY_ZOOMRESET is the Application Control Zoom key.
	KEY_ZOOMRESET = 0x1a4

	// KEY_WORDPROCESSOR is the Application Launch Word Processor key.
	KEY_WORDPROCESSOR = 0x1a5

	// KEY_EDITOR is the Application Launch Text Editor key.
	KEY_EDITOR = 0x1a6

	// KEY_SPREADSHEET is the Application Launch Spreadsheet key.
	KEY_SPREADSHEET = 0x1a7

	// KEY_GRAPHICSEDITOR is the Application Launch Graphics Editor key.
	KEY_GRAPHICSEDITOR = 0x1a8

	// KEY_PRESENTATION is the Application Launch Presentation App key.
	KEY_PRESENTATION = 0x1a9

	// KEY_DATABASE is the Application Launch Database App key.
	KEY_DATABASE = 0x1aa

	// KEY_NEWS is the Application Launch Newsreader key.
	KEY_NEWS = 0x1ab

	// KEY_VOICEMAIL is the Application Launch Voicemail key.
	KEY_VOICEMAIL = 0x1ac

	// KEY_ADDRESSBOOK is the Application Launch Contacts/Address Book key.
	KEY_ADDRESSBOOK = 0x1ad

	// KEY_MESSENGER is the Application Launch Instant Messaging key.
	KEY_MESSENGER = 0x1ae

	// KEY_DISPLAYTOGGLE is the Turn display (LCD) on and off key.
	KEY_DISPLAYTOGGLE = 0x1af

	// KEY_BRIGHTNESS_TOGGLE is an alias for KEY_DISPLAYTOGGLE.
	KEY_BRIGHTNESS_TOGGLE = KEY_DISPLAYTOGGLE

	// KEY_SPELLCHECK is the Application Launch Spell Check key.
	KEY_SPELLCHECK = 0x1b0

	// KEY_LOGOFF is the Application Launch Logoff key.
	KEY_LOGOFF = 0x1b1

	// KEY_DOLLAR is the dollar sign ($) key code.
	KEY_DOLLAR = 0x1b2

	// KEY_EURO is the euro sign (€) key code.
	KEY_EURO = 0x1b3

	// KEY_FRAMEBACK is the frame-backward transport control key code.
	KEY_FRAMEBACK = 0x1b4

	// KEY_FRAMEFORWARD is the frame-forward transport control key code.
	KEY_FRAMEFORWARD = 0x1b5

	// KEY_CONTEXT_MENU is the system context menu key code.
	KEY_CONTEXT_MENU = 0x1b6

	// KEY_MEDIA_REPEAT is the media repeat transport control key code.
	KEY_MEDIA_REPEAT = 0x1b7

	// KEY_10CHANNELSUP is the ten-channels-up key code.
	KEY_10CHANNELSUP = 0x1b8

	// KEY_10CHANNELSDOWN is the ten-channels-down key code.
	KEY_10CHANNELSDOWN = 0x1b9

	// KEY_IMAGES is the image browser key code.
	KEY_IMAGES = 0x1ba

	// KEY_NOTIFICATION_CENTER is the notification center toggle key code.
	KEY_NOTIFICATION_CENTER = 0x1bc

	// KEY_PICKUP_PHONE is the answer incoming call key code.
	KEY_PICKUP_PHONE = 0x1bd

	// KEY_HANGUP_PHONE is the decline incoming call key code.
	KEY_HANGUP_PHONE = 0x1be

	// KEY_LINK_PHONE is the phone sync key code.
	KEY_LINK_PHONE = 0x1bf

	// KEY_DEL_EOL deletes text from the cursor to the end of the line.
	KEY_DEL_EOL = 0x1c0

	// KEY_DEL_EOS deletes text from the cursor to the end of the screen.
	KEY_DEL_EOS = 0x1c1

	// KEY_INS_LINE inserts a new line at the cursor position.
	KEY_INS_LINE = 0x1c2

	// KEY_DEL_LINE deletes the entire current line.
	KEY_DEL_LINE = 0x1c3

	// KEY_FN is the function (Fn) modifier key.
	KEY_FN = 0x1d0

	// KEY_FN_ESC is the Fn+Esc key.
	KEY_FN_ESC = 0x1d1

	// KEY_FN_F1 is the Fn+F1 key.
	KEY_FN_F1 = 0x1d2

	// KEY_FN_F2 is the Fn+F2 key.
	KEY_FN_F2 = 0x1d3

	// KEY_FN_F3 is the Fn+F3 key.
	KEY_FN_F3 = 0x1d4

	// KEY_FN_F4 is the Fn+F4 key.
	KEY_FN_F4 = 0x1d5

	// KEY_FN_F5 is the Fn+F5 key.
	KEY_FN_F5 = 0x1d6

	// KEY_FN_F6 is the Fn+F6 key.
	KEY_FN_F6 = 0x1d7

	// KEY_FN_F7 is the Fn+F7 key.
	KEY_FN_F7 = 0x1d8

	// KEY_FN_F8 is the Fn+F8 key.
	KEY_FN_F8 = 0x1d9

	// KEY_FN_F9 is the Fn+F9 key.
	KEY_FN_F9 = 0x1da

	// KEY_FN_F10 is the Fn+F10 key.
	KEY_FN_F10 = 0x1db

	// KEY_FN_F11 is the Fn+F11 key.
	KEY_FN_F11 = 0x1dc

	// KEY_FN_F12 is the Fn+F12 key.
	KEY_FN_F12 = 0x1dd

	// KEY_FN_1 is the Fn+1 key.
	KEY_FN_1 = 0x1de

	// KEY_FN_2 is the Fn+2 key.
	KEY_FN_2 = 0x1df

	// KEY_FN_D is the Fn+D key.
	KEY_FN_D = 0x1e0

	// KEY_FN_E is the Fn+E key.
	KEY_FN_E = 0x1e1

	// KEY_FN_F is the Fn+F key.
	KEY_FN_F = 0x1e2

	// KEY_FN_S is the Fn+S key.
	KEY_FN_S = 0x1e3

	// KEY_FN_B is the Fn+B key.
	KEY_FN_B = 0x1e4

	// KEY_FN_RIGHT_SHIFT is the Fn+Right Shift key.
	KEY_FN_RIGHT_SHIFT = 0x1e5

	// KEY_BRL_DOT1 is the Braille dot 1 key code.
	KEY_BRL_DOT1 = 0x1f1

	// KEY_BRL_DOT2 is the Braille dot 2 key code.
	KEY_BRL_DOT2 = 0x1f2

	// KEY_BRL_DOT3 is the Braille dot 3 key code.
	KEY_BRL_DOT3 = 0x1f3

	// KEY_BRL_DOT4 is the Braille dot 4 key code.
	KEY_BRL_DOT4 = 0x1f4

	// KEY_BRL_DOT5 is the Braille dot 5 key code.
	KEY_BRL_DOT5 = 0x1f5

	// KEY_BRL_DOT6 is the Braille dot 6 key code.
	KEY_BRL_DOT6 = 0x1f6

	// KEY_BRL_DOT7 is the Braille dot 7 key code.
	KEY_BRL_DOT7 = 0x1f7

	// KEY_BRL_DOT8 is the Braille dot 8 key code.
	KEY_BRL_DOT8 = 0x1f8

	// KEY_BRL_DOT9 is the Braille dot 9 key code.
	KEY_BRL_DOT9 = 0x1f9

	// KEY_BRL_DOT10 is the Braille dot 10 key code.
	KEY_BRL_DOT10 = 0x1fa

	// KEY_NUMERIC_0 is the 0 key on a phone or remote control keypad.
	KEY_NUMERIC_0 = 0x200

	// KEY_NUMERIC_1 is the 1 key on a phone or remote control keypad.
	KEY_NUMERIC_1 = 0x201

	// KEY_NUMERIC_2 is the 2 key on a phone or remote control keypad.
	KEY_NUMERIC_2 = 0x202

	// KEY_NUMERIC_3 is the 3 key on a phone or remote control keypad.
	KEY_NUMERIC_3 = 0x203

	// KEY_NUMERIC_4 is the 4 key on a phone or remote control keypad.
	KEY_NUMERIC_4 = 0x204

	// KEY_NUMERIC_5 is the 5 key on a phone or remote control keypad.
	KEY_NUMERIC_5 = 0x205

	// KEY_NUMERIC_6 is the 6 key on a phone or remote control keypad.
	KEY_NUMERIC_6 = 0x206

	// KEY_NUMERIC_7 is the 7 key on a phone or remote control keypad.
	KEY_NUMERIC_7 = 0x207

	// KEY_NUMERIC_8 is the 8 key on a phone or remote control keypad.
	KEY_NUMERIC_8 = 0x208

	// KEY_NUMERIC_9 is the 9 key on a phone or remote control keypad.
	KEY_NUMERIC_9 = 0x209

	// KEY_NUMERIC_STAR is the star (*) key on a phone keypad.
	KEY_NUMERIC_STAR = 0x20a

	// KEY_NUMERIC_POUND is the pound (#) key on a phone keypad.
	KEY_NUMERIC_POUND = 0x20b

	// KEY_NUMERIC_A is the A key on a phone keypad.
	KEY_NUMERIC_A = 0x20c

	// KEY_NUMERIC_B is the B key on a phone keypad.
	KEY_NUMERIC_B = 0x20d

	// KEY_NUMERIC_C is the C key on a phone keypad.
	KEY_NUMERIC_C = 0x20e

	// KEY_NUMERIC_D is the D key on a phone keypad.
	KEY_NUMERIC_D = 0x20f

	// KEY_CAMERA_FOCUS is the camera focus key code.
	KEY_CAMERA_FOCUS = 0x210

	// KEY_WPS_BUTTON is the Wi-Fi Protected Setup button key code.
	KEY_WPS_BUTTON = 0x211

	// KEY_TOUCHPAD_TOGGLE toggles the touchpad on or off.
	KEY_TOUCHPAD_TOGGLE = 0x212

	// KEY_TOUCHPAD_ON turns the touchpad on.
	KEY_TOUCHPAD_ON = 0x213

	// KEY_TOUCHPAD_OFF turns the touchpad off.
	KEY_TOUCHPAD_OFF = 0x214

	// KEY_CAMERA_ZOOMIN zooms the camera in.
	KEY_CAMERA_ZOOMIN = 0x215

	// KEY_CAMERA_ZOOMOUT zooms the camera out.
	KEY_CAMERA_ZOOMOUT = 0x216

	// KEY_CAMERA_UP moves the camera view up.
	KEY_CAMERA_UP = 0x217

	// KEY_CAMERA_DOWN moves the camera view down.
	KEY_CAMERA_DOWN = 0x218

	// KEY_CAMERA_LEFT moves the camera view left.
	KEY_CAMERA_LEFT = 0x219

	// KEY_CAMERA_RIGHT moves the camera view right.
	KEY_CAMERA_RIGHT = 0x21a

	// KEY_ATTENDANT_ON signals attendant call on.
	KEY_ATTENDANT_ON = 0x21b

	// KEY_ATTENDANT_OFF signals attendant call off.
	KEY_ATTENDANT_OFF = 0x21c

	// KEY_ATTENDANT_TOGGLE toggles attendant call state.
	KEY_ATTENDANT_TOGGLE = 0x21d

	// KEY_LIGHTS_TOGGLE toggles the reading light on or off.
	KEY_LIGHTS_TOGGLE = 0x21e

	// BTN_DPAD_UP is the directional pad up button code.
	BTN_DPAD_UP = 0x220

	// BTN_DPAD_DOWN is the directional pad down button code.
	BTN_DPAD_DOWN = 0x221

	// BTN_DPAD_LEFT is the directional pad left button code.
	BTN_DPAD_LEFT = 0x222

	// BTN_DPAD_RIGHT is the directional pad right button code.
	BTN_DPAD_RIGHT = 0x223

	// KEY_ALS_TOGGLE toggles the ambient light sensor.
	KEY_ALS_TOGGLE = 0x230

	// KEY_ROTATE_LOCK_TOGGLE toggles screen rotation lock.
	KEY_ROTATE_LOCK_TOGGLE = 0x231

	// KEY_REFRESH_RATE_TOGGLE toggles display refresh rate.
	KEY_REFRESH_RATE_TOGGLE = 0x232

	// KEY_BUTTONCONFIG is the application launch button configuration key.
	KEY_BUTTONCONFIG = 0x240

	// KEY_TASKMANAGER is the application launch task manager key.
	KEY_TASKMANAGER = 0x241

	// KEY_JOURNAL is the application launch log/journal key.
	KEY_JOURNAL = 0x242

	// KEY_CONTROLPANEL is the application launch control panel key.
	KEY_CONTROLPANEL = 0x243

	// KEY_APPSELECT is the application launch app selection key.
	KEY_APPSELECT = 0x244

	// KEY_SCREENSAVER is the application launch screen saver key.
	KEY_SCREENSAVER = 0x245

	// KEY_VOICECOMMAND is the voice command activation key.
	KEY_VOICECOMMAND = 0x246

	// KEY_ASSISTANT is the context-aware assistant activation key.
	KEY_ASSISTANT = 0x247

	// KEY_KBD_LAYOUT_NEXT selects the next keyboard layout.
	KEY_KBD_LAYOUT_NEXT = 0x248

	// KEY_EMOJI_PICKER shows or hides the emoji picker.
	KEY_EMOJI_PICKER = 0x249

	// KEY_DICTATE starts or stops voice dictation.
	KEY_DICTATE = 0x24a

	// KEY_CAMERA_ACCESS_ENABLE enables programmatic camera access.
	KEY_CAMERA_ACCESS_ENABLE = 0x24b

	// KEY_CAMERA_ACCESS_DISABLE disables programmatic camera access.
	KEY_CAMERA_ACCESS_DISABLE = 0x24c

	// KEY_CAMERA_ACCESS_TOGGLE toggles programmatic camera access.
	KEY_CAMERA_ACCESS_TOGGLE = 0x24d

	// KEY_ACCESSIBILITY toggles the system accessibility UI.
	KEY_ACCESSIBILITY = 0x24e

	// KEY_DO_NOT_DISTURB toggles Do Not Disturb mode.
	KEY_DO_NOT_DISTURB = 0x24f

	// KEY_BRIGHTNESS_MIN sets brightness to minimum.
	KEY_BRIGHTNESS_MIN = 0x250

	// KEY_BRIGHTNESS_MAX sets brightness to maximum.
	KEY_BRIGHTNESS_MAX = 0x251

	// KEY_KBDINPUTASSIST_PREV selects the previous input suggestion.
	KEY_KBDINPUTASSIST_PREV = 0x260

	// KEY_KBDINPUTASSIST_NEXT selects the next input suggestion.
	KEY_KBDINPUTASSIST_NEXT = 0x261

	// KEY_KBDINPUTASSIST_PREVGROUP moves to the previous suggestion group.
	KEY_KBDINPUTASSIST_PREVGROUP = 0x262

	// KEY_KBDINPUTASSIST_NEXTGROUP moves to the next suggestion group.
	KEY_KBDINPUTASSIST_NEXTGROUP = 0x263

	// KEY_KBDINPUTASSIST_ACCEPT accepts the current input suggestion.
	KEY_KBDINPUTASSIST_ACCEPT = 0x264

	// KEY_KBDINPUTASSIST_CANCEL cancels the current input suggestion.
	KEY_KBDINPUTASSIST_CANCEL = 0x265

	// KEY_RIGHT_UP is the diagonal up-right navigation key.
	KEY_RIGHT_UP = 0x266

	// KEY_RIGHT_DOWN is the diagonal down-right navigation key.
	KEY_RIGHT_DOWN = 0x267

	// KEY_LEFT_UP is the diagonal up-left navigation key.
	KEY_LEFT_UP = 0x268

	// KEY_LEFT_DOWN is the diagonal down-left navigation key.
	KEY_LEFT_DOWN = 0x269

	// KEY_ROOT_MENU shows the device’s root menu.
	KEY_ROOT_MENU = 0x26a

	// KEY_MEDIA_TOP_MENU shows the top menu of media (e.g. DVD).
	KEY_MEDIA_TOP_MENU = 0x26b

	// KEY_NUMERIC_11 is the 11 key on a phone or remote-control keypad.
	KEY_NUMERIC_11 = 0x26c

	// KEY_NUMERIC_12 is the 12 key on a phone or remote-control keypad.
	KEY_NUMERIC_12 = 0x26d

	// KEY_AUDIO_DESC toggles audio description for visually impaired users.
	KEY_AUDIO_DESC = 0x26e

	// KEY_3D_MODE toggles 3D display mode.
	KEY_3D_MODE = 0x26f

	// KEY_NEXT_FAVORITE goes to the next favorite channel.
	KEY_NEXT_FAVORITE = 0x270

	// KEY_STOP_RECORD stops recording.
	KEY_STOP_RECORD = 0x271

	// KEY_PAUSE_RECORD pauses recording.
	KEY_PAUSE_RECORD = 0x272

	// KEY_VOD launches video on demand.
	KEY_VOD = 0x273

	// KEY_UNMUTE unmutes audio.
	KEY_UNMUTE = 0x274

	// KEY_FASTREVERSE plays content in fast reverse.
	KEY_FASTREVERSE = 0x275

	// KEY_SLOWREVERSE plays content in slow reverse.
	KEY_SLOWREVERSE = 0x276

	// KEY_DATA controls interactive data applications on the current
	// channel.
	KEY_DATA = 0x277

	// KEY_ONSCREEN_KEYBOARD toggles the on-screen keyboard.
	KEY_ONSCREEN_KEYBOARD = 0x278

	// KEY_PRIVACY_SCREEN_TOGGLE toggles the electronic privacy screen.
	KEY_PRIVACY_SCREEN_TOGGLE = 0x279

	// KEY_SELECTIVE_SCREENSHOT captures a selected area of the screen.
	KEY_SELECTIVE_SCREENSHOT = 0x27a

	// KEY_NEXT_ELEMENT moves focus to the next element in the user
	// interface.
	KEY_NEXT_ELEMENT = 0x27b

	// KEY_PREVIOUS_ELEMENT moves focus to the previous element in the user
	// interface.
	KEY_PREVIOUS_ELEMENT = 0x27c

	// KEY_AUTOPILOT_ENGAGE_TOGGLE toggles autopilot engagement.
	KEY_AUTOPILOT_ENGAGE_TOGGLE = 0x27d

	// KEY_MARK_WAYPOINT marks the current position as a waypoint.
	KEY_MARK_WAYPOINT = 0x27e

	// KEY_SOS sends an SOS distress signal.
	KEY_SOS = 0x27f

	// KEY_NAV_CHART shows the navigation chart.
	KEY_NAV_CHART = 0x280

	// KEY_FISHING_CHART shows the fishing chart.
	KEY_FISHING_CHART = 0x281

	// KEY_SINGLE_RANGE_RADAR activates single-range radar.
	KEY_SINGLE_RANGE_RADAR = 0x282

	// KEY_DUAL_RANGE_RADAR activates dual-range radar.
	KEY_DUAL_RANGE_RADAR = 0x283

	// KEY_RADAR_OVERLAY toggles the radar overlay.
	KEY_RADAR_OVERLAY = 0x284

	// KEY_TRADITIONAL_SONAR activates traditional sonar.
	KEY_TRADITIONAL_SONAR = 0x285

	// KEY_CLEARVU_SONAR activates ClearVu down-imaging sonar.
	KEY_CLEARVU_SONAR = 0x286

	// KEY_SIDEVU_SONAR activates SideVu side-imaging sonar.
	KEY_SIDEVU_SONAR = 0x287

	// KEY_NAV_INFO shows navigation information.
	KEY_NAV_INFO = 0x288

	// KEY_BRIGHTNESS_MENU opens the brightness settings menu.
	KEY_BRIGHTNESS_MENU = 0x289

	// KEY_MACRO1 is a user-programmable macro key.
	KEY_MACRO1 = 0x290

	// KEY_MACRO2 is a user-programmable macro key.
	KEY_MACRO2 = 0x291

	// KEY_MACRO3 is a user-programmable macro key.
	KEY_MACRO3 = 0x292

	// KEY_MACRO4 is a user-programmable macro key.
	KEY_MACRO4 = 0x293

	// KEY_MACRO5 is a user-programmable macro key.
	KEY_MACRO5 = 0x294

	// KEY_MACRO6 is a user-programmable macro key.
	KEY_MACRO6 = 0x295

	// KEY_MACRO7 is a user-programmable macro key.
	KEY_MACRO7 = 0x296

	// KEY_MACRO8 is a user-programmable macro key.
	KEY_MACRO8 = 0x297

	// KEY_MACRO9 is a user-programmable macro key.
	KEY_MACRO9 = 0x298

	// KEY_MACRO10 is a user-programmable macro key.
	KEY_MACRO10 = 0x299

	// KEY_MACRO11 is a user-programmable macro key.
	KEY_MACRO11 = 0x29a

	// KEY_MACRO12 is a user-programmable macro key.
	KEY_MACRO12 = 0x29b

	// KEY_MACRO13 is a user-programmable macro key.
	KEY_MACRO13 = 0x29c

	// KEY_MACRO14 is a user-programmable macro key.
	KEY_MACRO14 = 0x29d

	// KEY_MACRO15 is a user-programmable macro key.
	KEY_MACRO15 = 0x29e

	// KEY_MACRO16 is a user-programmable macro key.
	KEY_MACRO16 = 0x29f

	// KEY_MACRO17 is a user-programmable macro key.
	KEY_MACRO17 = 0x2a0

	// KEY_MACRO18 is a user-programmable macro key.
	KEY_MACRO18 = 0x2a1

	// KEY_MACRO19 is a user-programmable macro key.
	KEY_MACRO19 = 0x2a2

	// KEY_MACRO20 is a user-programmable macro key.
	KEY_MACRO20 = 0x2a3

	// KEY_MACRO21 is a user-programmable macro key.
	KEY_MACRO21 = 0x2a4

	// KEY_MACRO22 is a user-programmable macro key.
	KEY_MACRO22 = 0x2a5

	// KEY_MACRO23 is a user-programmable macro key.
	KEY_MACRO23 = 0x2a6

	// KEY_MACRO24 is a user-programmable macro key.
	KEY_MACRO24 = 0x2a7

	// KEY_MACRO25 is a user-programmable macro key.
	KEY_MACRO25 = 0x2a8

	// KEY_MACRO26 is a user-programmable macro key.
	KEY_MACRO26 = 0x2a9

	// KEY_MACRO27 is a user-programmable macro key.
	KEY_MACRO27 = 0x2aa

	// KEY_MACRO28 is a user-programmable macro key.
	KEY_MACRO28 = 0x2ab

	// KEY_MACRO29 is a user-programmable macro key.
	KEY_MACRO29 = 0x2ac

	// KEY_MACRO30 is a user-programmable macro key.
	KEY_MACRO30 = 0x2ad

	// KEY_MACRO_RECORD_START starts macro recording.
	KEY_MACRO_RECORD_START = 0x2b0

	// KEY_MACRO_RECORD_STOP stops macro recording.
	KEY_MACRO_RECORD_STOP = 0x2b1

	// KEY_MACRO_PRESET_CYCLE cycles through macro presets.
	KEY_MACRO_PRESET_CYCLE = 0x2b2

	// KEY_MACRO_PRESET1 selects macro preset 1.
	KEY_MACRO_PRESET1 = 0x2b3

	// KEY_MACRO_PRESET2 selects macro preset 2.
	KEY_MACRO_PRESET2 = 0x2b4

	// KEY_MACRO_PRESET3 selects macro preset 3.
	KEY_MACRO_PRESET3 = 0x2b5

	// KEY_KBD_LCD_MENU1 is the first unlabeled LCD menu key.
	KEY_KBD_LCD_MENU1 = 0x2b8

	// KEY_KBD_LCD_MENU2 is the second unlabeled LCD menu key.
	KEY_KBD_LCD_MENU2 = 0x2b9

	// KEY_KBD_LCD_MENU3 is the third unlabeled LCD menu key.
	KEY_KBD_LCD_MENU3 = 0x2ba

	// KEY_KBD_LCD_MENU4 is the fourth unlabeled LCD menu key.
	KEY_KBD_LCD_MENU4 = 0x2bb

	// KEY_KBD_LCD_MENU5 is the fifth unlabeled LCD menu key.
	KEY_KBD_LCD_MENU5 = 0x2bc

	// BTN_TRIGGER_HAPPY is the first generic extra button code.
	BTN_TRIGGER_HAPPY = 0x2c0

	// BTN_TRIGGER_HAPPY1 is the first generic extra button code.
	BTN_TRIGGER_HAPPY1 = 0x2c0

	// BTN_TRIGGER_HAPPY2 is the second generic extra button code.
	BTN_TRIGGER_HAPPY2 = 0x2c1

	// BTN_TRIGGER_HAPPY3 is the third generic extra button code.
	BTN_TRIGGER_HAPPY3 = 0x2c2

	// BTN_TRIGGER_HAPPY4 is the fourth generic extra button code.
	BTN_TRIGGER_HAPPY4 = 0x2c3

	// BTN_TRIGGER_HAPPY5 is the fifth generic extra button code.
	BTN_TRIGGER_HAPPY5 = 0x2c4

	// BTN_TRIGGER_HAPPY6 is the sixth generic extra button code.
	BTN_TRIGGER_HAPPY6 = 0x2c5

	// BTN_TRIGGER_HAPPY7 is the seventh generic extra button code.
	BTN_TRIGGER_HAPPY7 = 0x2c6

	// BTN_TRIGGER_HAPPY8 is the eighth generic extra button code.
	BTN_TRIGGER_HAPPY8 = 0x2c7

	// BTN_TRIGGER_HAPPY9 is the ninth generic extra button code.
	BTN_TRIGGER_HAPPY9 = 0x2c8

	// BTN_TRIGGER_HAPPY10 is the tenth generic extra button code.
	BTN_TRIGGER_HAPPY10 = 0x2c9

	// BTN_TRIGGER_HAPPY11 is the eleventh generic extra button code.
	BTN_TRIGGER_HAPPY11 = 0x2ca

	// BTN_TRIGGER_HAPPY12 is the twelfth generic extra button code.
	BTN_TRIGGER_HAPPY12 = 0x2cb

	// BTN_TRIGGER_HAPPY13 is the thirteenth generic extra button code.
	BTN_TRIGGER_HAPPY13 = 0x2cc

	// BTN_TRIGGER_HAPPY14 is the fourteenth generic extra button code.
	BTN_TRIGGER_HAPPY14 = 0x2cd

	// BTN_TRIGGER_HAPPY15 is the fifteenth generic extra button code.
	BTN_TRIGGER_HAPPY15 = 0x2ce

	// BTN_TRIGGER_HAPPY16 is the sixteenth generic extra button code.
	BTN_TRIGGER_HAPPY16 = 0x2cf

	// BTN_TRIGGER_HAPPY17 is the seventeenth generic extra button code.
	BTN_TRIGGER_HAPPY17 = 0x2d0

	// BTN_TRIGGER_HAPPY18 is the eighteenth generic extra button code.
	BTN_TRIGGER_HAPPY18 = 0x2d1

	// BTN_TRIGGER_HAPPY19 is the nineteenth generic extra button code.
	BTN_TRIGGER_HAPPY19 = 0x2d2

	// BTN_TRIGGER_HAPPY20 is the twentieth generic extra button code.
	BTN_TRIGGER_HAPPY20 = 0x2d3

	// BTN_TRIGGER_HAPPY21 is the twenty-first generic extra button code.
	BTN_TRIGGER_HAPPY21 = 0x2d4

	// BTN_TRIGGER_HAPPY22 is the twenty-second generic extra button code.
	BTN_TRIGGER_HAPPY22 = 0x2d5

	// BTN_TRIGGER_HAPPY23 is the twenty-third generic extra button code.
	BTN_TRIGGER_HAPPY23 = 0x2d6

	// BTN_TRIGGER_HAPPY24 is the twenty-fourth generic extra button code.
	BTN_TRIGGER_HAPPY24 = 0x2d7

	// BTN_TRIGGER_HAPPY25 is the twenty-fifth generic extra button code.
	BTN_TRIGGER_HAPPY25 = 0x2d8

	// BTN_TRIGGER_HAPPY26 is the twenty-sixth generic extra button code.
	BTN_TRIGGER_HAPPY26 = 0x2d9

	// BTN_TRIGGER_HAPPY27 is the twenty-seventh generic extra button code.
	BTN_TRIGGER_HAPPY27 = 0x2da

	// BTN_TRIGGER_HAPPY28 is the twenty-eighth generic extra button code.
	BTN_TRIGGER_HAPPY28 = 0x2db

	// BTN_TRIGGER_HAPPY29 is the twenty-ninth generic extra button code.
	BTN_TRIGGER_HAPPY29 = 0x2dc

	// BTN_TRIGGER_HAPPY30 is the thirtieth generic extra button code.
	BTN_TRIGGER_HAPPY30 = 0x2dd

	// BTN_TRIGGER_HAPPY31 is the thirty-first generic extra button code.
	BTN_TRIGGER_HAPPY31 = 0x2de

	// BTN_TRIGGER_HAPPY32 is the thirty-second generic extra button code.
	BTN_TRIGGER_HAPPY32 = 0x2df

	// BTN_TRIGGER_HAPPY33 is the thirty-third generic extra button code.
	BTN_TRIGGER_HAPPY33 = 0x2e0

	// BTN_TRIGGER_HAPPY34 is the thirty-fourth generic extra button code.
	BTN_TRIGGER_HAPPY34 = 0x2e1

	// BTN_TRIGGER_HAPPY35 is the thirty-fifth generic extra button code.
	BTN_TRIGGER_HAPPY35 = 0x2e2

	// BTN_TRIGGER_HAPPY36 is the thirty-sixth generic extra button code.
	BTN_TRIGGER_HAPPY36 = 0x2e3

	// BTN_TRIGGER_HAPPY37 is the thirty-seventh generic extra button code.
	BTN_TRIGGER_HAPPY37 = 0x2e4

	// BTN_TRIGGER_HAPPY38 is the thirty-eighth generic extra button code.
	BTN_TRIGGER_HAPPY38 = 0x2e5

	// BTN_TRIGGER_HAPPY39 is the thirty-ninth generic extra button code.
	BTN_TRIGGER_HAPPY39 = 0x2e6

	// BTN_TRIGGER_HAPPY40 is the fortieth generic extra button code.
	BTN_TRIGGER_HAPPY40 = 0x2e7

	// KEY_MIN_INTERESTING is the lowest interesting key code.
	KEY_MIN_INTERESTING = KEY_MUTE

	// KEY_MAX is the highest key code value.
	KEY_MAX = 0x2ff

	// KEY_CNT is the total number of key codes.
	KEY_CNT = KEY_MAX + 1

	// REL_X is relative movement along the X axis.
	REL_X = 0x00

	// REL_Y is relative movement along the Y axis.
	REL_Y = 0x01

	// REL_Z is relative movement along the Z axis.
	REL_Z = 0x02

	// REL_RX is relative rotation around the X axis.
	REL_RX = 0x03

	// REL_RY is relative rotation around the Y axis.
	REL_RY = 0x04

	// REL_RZ is relative rotation around the Z axis.
	REL_RZ = 0x05

	// REL_HWHEEL is relative horizontal wheel movement.
	REL_HWHEEL = 0x06

	// REL_DIAL is relative dial rotation.
	REL_DIAL = 0x07

	// REL_WHEEL is relative vertical wheel movement.
	REL_WHEEL = 0x08

	// REL_MISC is a miscellaneous relative axis.
	REL_MISC = 0x09

	// REL_RESERVED is reserved and should not be used by input drivers.
	REL_RESERVED = 0x0a

	// REL_WHEEL_HI_RES is the high-resolution vertical wheel axis.
	REL_WHEEL_HI_RES = 0x0b

	// REL_HWHEEL_HI_RES is the high-resolution horizontal wheel axis.
	REL_HWHEEL_HI_RES = 0x0c

	// REL_MAX is the highest relative axis code.
	REL_MAX = 0x0f

	// REL_CNT is the total number of relative axis codes.
	REL_CNT = REL_MAX + 1

	// ABS_X is the absolute position along the X axis.
	ABS_X = 0x00

	// ABS_Y is the absolute position along the Y axis.
	ABS_Y = 0x01

	// ABS_Z is the absolute position along the Z axis.
	ABS_Z = 0x02

	// ABS_RX is the absolute rotation around the X axis.
	ABS_RX = 0x03

	// ABS_RY is the absolute rotation around the Y axis.
	ABS_RY = 0x04

	// ABS_RZ is the absolute rotation around the Z axis.
	ABS_RZ = 0x05

	// ABS_THROTTLE is the throttle control axis.
	ABS_THROTTLE = 0x06

	// ABS_RUDDER is the rudder control axis.
	ABS_RUDDER = 0x07

	// ABS_WHEEL is the steering wheel control axis.
	ABS_WHEEL = 0x08

	// ABS_GAS is the gas pedal control axis.
	ABS_GAS = 0x09

	// ABS_BRAKE is the brake pedal control axis.
	ABS_BRAKE = 0x0a

	// ABS_HAT0X is the horizontal axis of hat switch 0.
	ABS_HAT0X = 0x10

	// ABS_HAT0Y is the vertical axis of hat switch 0.
	ABS_HAT0Y = 0x11

	// ABS_HAT1X is the horizontal axis of hat switch 1.
	ABS_HAT1X = 0x12

	// ABS_HAT1Y is the vertical axis of hat switch 1.
	ABS_HAT1Y = 0x13

	// ABS_HAT2X is the horizontal axis of hat switch 2.
	ABS_HAT2X = 0x14

	// ABS_HAT2Y is the vertical axis of hat switch 2.
	ABS_HAT2Y = 0x15

	// ABS_HAT3X is the horizontal axis of hat switch 3.
	ABS_HAT3X = 0x16

	// ABS_HAT3Y is the vertical axis of hat switch 3.
	ABS_HAT3Y = 0x17

	// ABS_PRESSURE is the pressure axis (e.g., stylus pressure).
	ABS_PRESSURE = 0x18

	// ABS_DISTANCE is the distance axis (e.g., stylus distance).
	ABS_DISTANCE = 0x19

	// ABS_TILT_X is the tilt angle around the X axis.
	ABS_TILT_X = 0x1a

	// ABS_TILT_Y is the tilt angle around the Y axis.
	ABS_TILT_Y = 0x1b

	// ABS_TOOL_WIDTH is the tool width axis (e.g., eraser width).
	ABS_TOOL_WIDTH = 0x1c

	// ABS_VOLUME is the absolute volume axis.
	ABS_VOLUME = 0x20

	// ABS_PROFILE is the profile axis.
	ABS_PROFILE = 0x21

	// ABS_MISC is a miscellaneous absolute axis.
	ABS_MISC = 0x28

	// ABS_RESERVED is a reserved axis code used to detect invalid events.
	ABS_RESERVED = 0x2e

	// ABS_MT_SLOT is the multi-touch slot being modified.
	ABS_MT_SLOT = 0x2f

	// ABS_MT_TOUCH_MAJOR is the major axis of the touch ellipse.
	ABS_MT_TOUCH_MAJOR = 0x30

	// ABS_MT_TOUCH_MINOR is the minor axis of the touch ellipse.
	ABS_MT_TOUCH_MINOR = 0x31

	// ABS_MT_WIDTH_MAJOR is the major axis of the approaching ellipse.
	ABS_MT_WIDTH_MAJOR = 0x32

	// ABS_MT_WIDTH_MINOR is the minor axis of the approaching ellipse.
	ABS_MT_WIDTH_MINOR = 0x33

	// ABS_MT_ORIENTATION is the orientation of the touch ellipse.
	ABS_MT_ORIENTATION = 0x34

	// ABS_MT_POSITION_X is the X coordinate of the touch position.
	ABS_MT_POSITION_X = 0x35

	// ABS_MT_POSITION_Y is the Y coordinate of the touch position.
	ABS_MT_POSITION_Y = 0x36

	// ABS_MT_TOOL_TYPE is the type of tool in contact
	// (e.g., finger or stylus).
	ABS_MT_TOOL_TYPE = 0x37

	// ABS_MT_BLOB_ID groups packets into a single blob.
	ABS_MT_BLOB_ID = 0x38

	// ABS_MT_TRACKING_ID is a unique ID for touch contacts.
	ABS_MT_TRACKING_ID = 0x39

	// ABS_MT_PRESSURE is the pressure of the touch.
	ABS_MT_PRESSURE = 0x3a

	// ABS_MT_DISTANCE is the hover distance for touch.
	ABS_MT_DISTANCE = 0x3b

	// ABS_MT_TOOL_X is the X coordinate of the tool position.
	ABS_MT_TOOL_X = 0x3c

	// ABS_MT_TOOL_Y is the Y coordinate of the tool position.
	ABS_MT_TOOL_Y = 0x3d

	// ABS_MAX is the highest absolute axis code.
	ABS_MAX = 0x3f

	// ABS_CNT is the total number of absolute axis codes.
	ABS_CNT = ABS_MAX + 1

	// SW_LID indicates the lid is closed.
	SW_LID = 0x00

	// SW_TABLET_MODE indicates tablet mode is active.
	SW_TABLET_MODE = 0x01

	// SW_HEADPHONE_INSERT indicates headphones are inserted.
	SW_HEADPHONE_INSERT = 0x02

	// SW_RFKILL_ALL is the RF kill master switch (radio enabled).
	SW_RFKILL_ALL = 0x03

	// SW_RADIO is a deprecated alias for SW_RFKILL_ALL.
	SW_RADIO = SW_RFKILL_ALL

	// SW_MICROPHONE_INSERT indicates a microphone is inserted.
	SW_MICROPHONE_INSERT = 0x04

	// SW_DOCK indicates the device is docked.
	SW_DOCK = 0x05

	// SW_LINEOUT_INSERT indicates a line-out jack is connected.
	SW_LINEOUT_INSERT = 0x06

	// SW_JACK_PHYSICAL_INSERT indicates a mechanical jack is engaged.
	SW_JACK_PHYSICAL_INSERT = 0x07

	// SW_VIDEOOUT_INSERT indicates a video-out connector is attached.
	SW_VIDEOOUT_INSERT = 0x08

	// SW_CAMERA_LENS_COVER indicates the camera lens cover is down.
	SW_CAMERA_LENS_COVER = 0x09

	// SW_KEYPAD_SLIDE indicates the keypad is slid out.
	SW_KEYPAD_SLIDE = 0x0a

	// SW_FRONT_PROXIMITY indicates the front proximity sensor is active.
	SW_FRONT_PROXIMITY = 0x0b

	// SW_ROTATE_LOCK indicates screen rotation is locked.
	SW_ROTATE_LOCK = 0x0c

	// SW_LINEIN_INSERT indicates a line-in jack is connected.
	SW_LINEIN_INSERT = 0x0d

	// SW_MUTE_DEVICE indicates the device is muted.
	SW_MUTE_DEVICE = 0x0e

	// SW_PEN_INSERTED indicates a pen is inserted.
	SW_PEN_INSERTED = 0x0f

	// SW_MACHINE_COVER indicates the machine cover is closed.
	SW_MACHINE_COVER = 0x10

	// SW_USB_INSERT indicates a USB audio device is connected.
	SW_USB_INSERT = 0x11

	// SW_MAX is the highest switch event code.
	SW_MAX = 0x11

	// SW_CNT is the total number of switch event codes.
	SW_CNT = SW_MAX + 1

	// MSC_SERIAL is a serial event.
	MSC_SERIAL = 0x00

	// MSC_PULSELED is an LED pulse event.
	MSC_PULSELED = 0x01

	// MSC_GESTURE is a gesture event.
	MSC_GESTURE = 0x02

	// MSC_RAW is a raw data event.
	MSC_RAW = 0x03

	// MSC_SCAN is a scan code event.
	MSC_SCAN = 0x04

	// MSC_TIMESTAMP is a timestamp event.
	MSC_TIMESTAMP = 0x05

	// MSC_MAX is the highest miscellaneous event code.
	MSC_MAX = 0x07

	// MSC_CNT is the total number of miscellaneous event codes.
	MSC_CNT = MSC_MAX + 1

	// LED_NUML is the Num Lock LED.
	LED_NUML = 0x00

	// LED_CAPSL is the Caps Lock LED.
	LED_CAPSL = 0x01

	// LED_SCROLLL is the Scroll Lock LED.
	LED_SCROLLL = 0x02

	// LED_COMPOSE is the Compose LED.
	LED_COMPOSE = 0x03

	// LED_KANA is the Kana (input mode) LED.
	LED_KANA = 0x04

	// LED_SLEEP is the Sleep state LED.
	LED_SLEEP = 0x05

	// LED_SUSPEND is the Suspend state LED.
	LED_SUSPEND = 0x06

	// LED_MUTE is the Mute state LED.
	LED_MUTE = 0x07

	// LED_MISC is the miscellaneous LED.
	LED_MISC = 0x08

	// LED_MAIL is the Mail notification LED.
	LED_MAIL = 0x09

	// LED_CHARGING is the Charging state LED.
	LED_CHARGING = 0x0a

	// LED_MAX is the highest LED code.
	LED_MAX = 0x0f

	// LED_CNT is the total number of LED codes.
	LED_CNT = LED_MAX + 1

	// REP_DELAY is the autorepeat delay value.
	REP_DELAY = 0x00

	// REP_PERIOD is the autorepeat period value.
	REP_PERIOD = 0x01

	// REP_MAX is the highest autorepeat index.
	REP_MAX = 0x01

	// REP_CNT is the total number of autorepeat values.
	REP_CNT = REP_MAX + 1

	// SND_CLICK is the click sound code.
	SND_CLICK = 0x00

	// SND_BELL is the bell sound code.
	SND_BELL = 0x01

	// SND_TONE is the tone sound code.
	SND_TONE = 0x02

	// SND_MAX is the highest sound code.
	SND_MAX = 0x07

	// SND_CNT is the total number of sound codes.
	SND_CNT = SND_MAX + 1
)
