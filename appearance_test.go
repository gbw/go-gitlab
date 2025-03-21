//
// Copyright 2023, 徐晓伟 <xuxiaowei@xuxiaowei.com.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAppearance(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/application/appearance", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
 		 	"title": "GitLab Test Instance",
 		 	"description": "gitlab-test.example.com",
 		 	"pwa_name": "GitLab PWA",
 		 	"pwa_short_name": "GitLab",
 		 	"pwa_description": "GitLab as PWA",
 		 	"pwa_icon": "/uploads/-/system/appearance/pwa_icon/1/pwa_logo.png",
 		 	"logo": "/uploads/-/system/appearance/logo/1/logo.png",
 		 	"header_logo": "/uploads/-/system/appearance/header_logo/1/header.png",
 		 	"favicon": "/uploads/-/system/appearance/favicon/1/favicon.png",
 		 	"member_guidelines": "Please be nice.",
 		 	"new_project_guidelines": "Please read the FAQs for help.",
 		 	"profile_image_guidelines": "Custom profile image guidelines",
 		 	"header_message": "",
 		 	"footer_message": "",
 		 	"message_background_color": "#e75e40",
 		 	"message_font_color": "#ffffff",
 		 	"email_header_and_footer_enabled": false
 		}`)
	})

	appearance, _, err := client.Appearance.GetAppearance()
	if err != nil {
		t.Errorf("Appearance.GetAppearance returned error: %v", err)
	}

	want := &Appearance{
		Title:                       "GitLab Test Instance",
		Description:                 "gitlab-test.example.com",
		PWAName:                     "GitLab PWA",
		PWAShortName:                "GitLab",
		PWADescription:              "GitLab as PWA",
		PWAIcon:                     "/uploads/-/system/appearance/pwa_icon/1/pwa_logo.png",
		Logo:                        "/uploads/-/system/appearance/logo/1/logo.png",
		HeaderLogo:                  "/uploads/-/system/appearance/header_logo/1/header.png",
		Favicon:                     "/uploads/-/system/appearance/favicon/1/favicon.png",
		MemberGuidelines:            "Please be nice.",
		NewProjectGuidelines:        "Please read the FAQs for help.",
		ProfileImageGuidelines:      "Custom profile image guidelines",
		HeaderMessage:               "",
		FooterMessage:               "",
		MessageBackgroundColor:      "#e75e40",
		MessageFontColor:            "#ffffff",
		EmailHeaderAndFooterEnabled: false,
	}

	if !reflect.DeepEqual(want, appearance) {
		t.Errorf("Appearance.GetAppearance returned %+v, want %+v", appearance, want)
	}
}

func TestChangeAppearance(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/application/appearance", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{
		 	"title": "GitLab Test Instance - 001",
 		 	"description": "gitlab-test.example.com",
 		 	"pwa_name": "GitLab PWA",
 		 	"pwa_short_name": "GitLab",
 		 	"pwa_description": "GitLab as PWA",
 		 	"pwa_icon": "/uploads/-/system/appearance/pwa_icon/1/pwa_logo.png",
 		 	"logo": "/uploads/-/system/appearance/logo/1/logo.png",
 		 	"header_logo": "/uploads/-/system/appearance/header_logo/1/header.png",
 		 	"favicon": "/uploads/-/system/appearance/favicon/1/favicon.png",
 		 	"member_guidelines": "Please be nice.",
 		 	"new_project_guidelines": "Please read the FAQs for help.",
 		 	"profile_image_guidelines": "Custom profile image guidelines",
 		 	"header_message": "",
 		 	"footer_message": "",
 		 	"message_background_color": "#e75e40",
 		 	"message_font_color": "#ffffff",
 		 	"email_header_and_footer_enabled": false
 		}`)
	})

	opt := &ChangeAppearanceOptions{
		Title:                       Ptr("GitLab Test Instance - 001"),
		Description:                 Ptr("gitlab-test.example.com"),
		PWAName:                     Ptr("GitLab PWA"),
		PWAShortName:                Ptr("GitLab"),
		PWADescription:              Ptr("GitLab as PWA"),
		PWAIcon:                     Ptr("/uploads/-/system/appearance/pwa_icon/1/pwa_logo.png"),
		Logo:                        Ptr("/uploads/-/system/appearance/logo/1/logo.png"),
		HeaderLogo:                  Ptr("/uploads/-/system/appearance/header_logo/1/header.png"),
		Favicon:                     Ptr("/uploads/-/system/appearance/favicon/1/favicon.png"),
		MemberGuidelines:            Ptr("Please be nice."),
		NewProjectGuidelines:        Ptr("Please read the FAQs for help."),
		ProfileImageGuidelines:      Ptr("Custom profile image guidelines"),
		HeaderMessage:               Ptr(""),
		FooterMessage:               Ptr(""),
		MessageBackgroundColor:      Ptr("#e75e40"),
		MessageFontColor:            Ptr("#ffffff"),
		EmailHeaderAndFooterEnabled: Ptr(false),
	}

	appearance, _, err := client.Appearance.ChangeAppearance(opt)
	if err != nil {
		t.Errorf("Appearance.ChangeAppearance returned error: %v", err)
	}

	want := &Appearance{
		Title:                       "GitLab Test Instance - 001",
		Description:                 "gitlab-test.example.com",
		PWAName:                     "GitLab PWA",
		PWAShortName:                "GitLab",
		PWADescription:              "GitLab as PWA",
		PWAIcon:                     "/uploads/-/system/appearance/pwa_icon/1/pwa_logo.png",
		Logo:                        "/uploads/-/system/appearance/logo/1/logo.png",
		HeaderLogo:                  "/uploads/-/system/appearance/header_logo/1/header.png",
		Favicon:                     "/uploads/-/system/appearance/favicon/1/favicon.png",
		MemberGuidelines:            "Please be nice.",
		NewProjectGuidelines:        "Please read the FAQs for help.",
		ProfileImageGuidelines:      "Custom profile image guidelines",
		HeaderMessage:               "",
		FooterMessage:               "",
		MessageBackgroundColor:      "#e75e40",
		MessageFontColor:            "#ffffff",
		EmailHeaderAndFooterEnabled: false,
	}

	if !reflect.DeepEqual(want, appearance) {
		t.Errorf("Appearance.GetAppearance returned %+v, want %+v", appearance, want)
	}
}
