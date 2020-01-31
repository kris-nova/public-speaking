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
	"strings"

	"github.com/kris-nova/logger"

	"github.com/nsf/termbox-go"
)

type DisplayOptions struct {
}

func (r RenderedSlide) DisplayWithOptions(opt *DisplayOptions) (DisplayAction, error) {
	err := termbox.Init()
	if err != nil {
		return nil, fmt.Errorf("unable to init termbox: %v", err)
	}
	defer termbox.Close()

	err = termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		return nil, fmt.Errorf("unable to clear screen: %v", err)
	}

	// Calculate the longest message
	highestLen := 0
	for _, ec := range r.elements {
		l := len(ec.RawString())
		if len(ec.RawString()) > highestLen {
			highestLen = l
		}
	}
	r.elements = append(r.elements, r.slide.presentation.title)
	r.elements = append(r.elements, r.slide.presentation.author)
	r.elements = append(r.elements, r.slide.presentation.meta)

	for i, e := range r.elements {
		// T's are center aligned
		// nl's are centered but aligned
		switch e.Label() {
		case nl:
			err := r.DisplayAlignedMessage(e, i, highestLen)
			if err != nil {
				return nil, fmt.Errorf("failed to display slide: %v", err)
			}
			break
		case ex:
			// Execute a command
			execStr := strings.Trim(e.RawString(), " ")
			//logger.Info("exec: %s", execStr)
			execSlc := strings.Split(execStr, " ")
			cmd := exec.Command(execSlc[0], execSlc[1:]...)
			cmd.Stdin = os.Stdin
			out, err := cmd.Output()
			if err != nil {
				return nil, fmt.Errorf("unable to execute command %s %s: %v", execStr[0], execStr[1:], err)
			}
			err = r.DisplayAlignedMessage(e, i, highestLen)
			if err != nil {
				return nil, fmt.Errorf("failed to display slide: %v", err)
			}
			outSlc := strings.Split(string(out), "\n")
			// This is dangerous and could break the program
			for _, line := range outSlc {
				i++
				// Todo the output will need formatting and a count as well - just using this for now
				eee := &EXElement{
					rawLine: line,
					label:   ex,
				}
				err := r.DisplayAlignedMessage(eee, i, highestLen)
				if err != nil {
					return nil, fmt.Errorf("failed to display slide: %v", err)
				}
			}
			break
		case ll:
			err := r.DisplayLeftAlignedMessage(e, r.maxx-1)
			if err != nil {
				return nil, fmt.Errorf("failed to display slide: %v", err)
			}
			break
		case rr:
			err := r.DisplayRightAlignedMessage(e, r.maxx-1)
			if err != nil {
				return nil, fmt.Errorf("failed to display slide: %v", err)
			}
			break
		case tt:
			err := r.DisplayCenteredMessage(e, -1)
			if err != nil {
				return nil, fmt.Errorf("failed to display slide: %v", err)
			}
			break
		default:
			err := r.DisplayCenteredMessage(e, i)
			if err != nil {
				return nil, fmt.Errorf("failed to display slide: %v", err)
			}
		}
	}
	termbox.Flush()
	event := termbox.PollEvent()
	switch event.Type {
	case termbox.EventKey:
		logger.Info("event [%v] (%v)", event.Key, event.Type)
		switch event.Key {
		case termbox.KeyArrowLeft:
			return DisplayActionPrev{}, nil
		case termbox.KeyArrowRight:
			return DisplayActionNext{}, nil
		case termbox.KeyHome:
			return DisplayActionStart{}, nil
		case termbox.KeyEnd:
			return DisplayActionEnd{}, nil
		}
		break
	case termbox.EventNone:
		break
	}
	return DisplayActionNext{}, nil
}

func (r *RenderedSlide) DisplayAlignedMessage(e Element, lnum, longestMsg int) error {
	msg := e.RawString()
	// Time for some arithmetic!
	// Horizontal Axis
	yEmpty := r.maxy - longestMsg
	if yEmpty < 0 {
		return fmt.Errorf("slide %s too many horizontal chars (%d) max chars (%d)", r.slide.name, (yEmpty * -1), longestMsg)
	}
	yStart := yEmpty/2 + ((yEmpty % 2) / 2)
	// Vertical Axis
	totalVertI := len(r.elements)
	xEmpty := r.maxx - totalVertI - 2 // Reserve 2 lines for the title and author lines
	if xEmpty < 0 {
		return fmt.Errorf("slide %s too many vertical lines (%d) max lines (%d)", r.slide.name, (xEmpty * -1), totalVertI)
	}
	xStart := xEmpty/2 + (xEmpty % 2)
	theme := r.slide.presentation.theme
	a1, a2 := theme.AttrsByElement(e)
	for i := 0; i <= longestMsg; i++ {
		for i, r := range msg {
			termbox.SetCell(i+yStart, lnum+xStart, r, a1, a2)
			i++
		}
		termbox.SetCell(i+yStart, lnum+xStart, 32, a1, a2)
	}
	return nil
}

func (r *RenderedSlide) DisplayCenteredMessage(e Element, lnum int) error {
	// Time for some arithmetic!
	// Horizontal Axis
	msg := e.RawString()
	totalLineI := len(msg)
	yEmpty := r.maxy - totalLineI
	if yEmpty < 0 {
		return fmt.Errorf("slide %s too many horizontal chars (%d) max chars (%d)", r.slide.name, (yEmpty * -1), totalLineI)
	}
	yStart := yEmpty/2 + (yEmpty % 2)
	// Vertical Axis
	xStart := 0
	if lnum == -1 {
		xStart = 1
	} else {
		totalVertI := len(r.elements)
		xEmpty := r.maxx - totalVertI - 2 // Reserve 2 lines
		if xEmpty < 0 {
			return fmt.Errorf("slide %s too many vertical lines (%d) max lines (%d)", r.slide.name, (xEmpty * -1), totalVertI)
		}
		xStart = xEmpty/2 + (xEmpty % 2)
	}
	theme := r.slide.presentation.theme
	a1, a2 := theme.AttrsByElement(e)
	for i, r := range msg {
		termbox.SetCell(i+yStart, lnum+xStart, r, a1, a2)
	}
	return nil
}

func (r *RenderedSlide) DisplayLeftAlignedMessage(e Element, lnum int) error {
	// Time for some arithmetic!
	msg := e.RawString()
	theme := r.slide.presentation.theme
	a1, a2 := theme.AttrsByElement(e)
	for i, r := range msg {
		termbox.SetCell(i, lnum, r, a1, a2)
	}
	return nil
}

func (r *RenderedSlide) DisplayRightAlignedMessage(e Element, lnum int) error {
	// Time for some arithmetic!
	msg := e.RawString()
	theme := r.slide.presentation.theme
	yStart := r.maxy - len(msg)
	a1, a2 := theme.AttrsByElement(e)
	for i, r := range msg {
		termbox.SetCell(i+yStart, lnum, r, a1, a2)
	}
	return nil
}
