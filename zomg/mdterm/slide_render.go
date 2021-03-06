/*
Copyright © 2019 Kris Nova <kris@nivenly.com> 2019

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package mdterm

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type RenderedSlide struct {
	elements []Element
	slide    *Slide
	maxx     int
	maxy     int
}

func RenderSlide(slide *Slide) (*RenderedSlide, error) {
	r := &RenderedSlide{
		slide: slide,
	}
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("unable to calculate terminal size to render slide: %v", err)
	}
	out1line := strings.Replace(string(out), "\n", "", -1)
	ll := strings.Split(out1line, " ")
	if len(ll) != 2 {
		return nil, fmt.Errorf("unable to calculate terminal size: %v", err)
	}
	x, err := strconv.Atoi(ll[0])
	if err != nil {
		return nil, fmt.Errorf("unable to calculate x term size: %v", err)
	}
	y, err := strconv.Atoi(ll[1])
	if err != nil {
		return nil, fmt.Errorf("unable to calculate y term size: %v", err)
	}
	lines := strings.Split(string(slide.content), "\n")
	for _, line := range lines {
		// t1 ----------------------------------------------------------------------------------------------------------
		if strings.HasPrefix(line, string(t1)) {
			e := &T1Element{
				label:   t1,
				rawLine: strings.TrimPrefix(line, string(t1)),
			}
			r.elements = append(r.elements, e)
		}
		// t2 ----------------------------------------------------------------------------------------------------------
		if strings.HasPrefix(line, string(t2)) {
			e := &T2Element{
				label:   t2,
				rawLine: strings.TrimPrefix(line, string(t2)),
			}
			r.elements = append(r.elements, e)
		}
		// t3 ----------------------------------------------------------------------------------------------------------
		if strings.HasPrefix(line, string(t3)) {
			e := &T1Element{
				label:   t3,
				rawLine: strings.TrimPrefix(line, string(t3)),
			}
			r.elements = append(r.elements, e)
		}
		// t4 ----------------------------------------------------------------------------------------------------------
		if strings.HasPrefix(line, string(t4)) {
			e := &T1Element{
				label:   t4,
				rawLine: strings.TrimPrefix(line, string(t4)),
			}
			r.elements = append(r.elements, e)
		}
		// nl  ----------------------------------------------------------------------------------------------------------
		if strings.HasPrefix(line, string(nl)) {
			e := &NLElement{
				label:   nl,
				rawLine: strings.TrimPrefix(line, string(nl)),
			}
			r.elements = append(r.elements, e)
		}
		// x  ----------------------------------------------------------------------------------------------------------
		if strings.HasPrefix(line, string(ex)) {
			e := &EXElement{
				label:   ex,
				rawLine: strings.TrimPrefix(line, string(ex)),
			}
			r.elements = append(r.elements, e)
		}
		// c  ----------------------------------------------------------------------------------------------------------
		if strings.HasPrefix(line, string(c)) {
			e := &NLElement{
				label:   c,
				rawLine: "",
			}
			r.elements = append(r.elements, e)
		}
	}

	// -----------------------------------------------------------------------------------------------------------------
	//
	// Validation
	//
	for _, e := range r.elements {
		l := len(e.RawString())
		if l > y {
			return nil, fmt.Errorf("word wrap not possible: invalid string length: %s", e.RawString())
		}
	}
	r.maxx = x
	r.maxy = y

	return r, nil
}
