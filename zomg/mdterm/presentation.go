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
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/kris-nova/logger"
)

const (
	ZomgExtension = ".zomg"
	ZomgREADMEmd  = "README.md"
)

type Presentation struct {
	filepath       string
	slides         []*Slide
	theme          Theme
	renderedSlides []*RenderedSlide
	title          Element
	author         Element
	meta           Element
}

type Slide struct {
	name         string
	content      []byte
	presentation *Presentation
}

func LoadPresentation(dir string) (*Presentation, error) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return nil, fmt.Errorf("unable to abs() dir: %v", dir)
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("unable to parse dir (%s): %v", dir, err)
	}
	p := &Presentation{}
	logger.Info("number of files in (%s): %d", dir, len(files))
	if len(files) < 1 {
		return nil, fmt.Errorf("unable to find enough files in (%s)", dir)
	}
	for _, f := range files {
		logger.Info("file: %s", f.Name())
		absf := filepath.Join(dir, f.Name())
		data, err := ioutil.ReadFile(absf)
		if err != nil {
			logger.Info("error parsing file: %v", err)
		}
		if f.Name() == ZomgREADMEmd {
			p.parseREADMEmd(string(data))
		}
		if !strings.HasSuffix(f.Name(), ZomgExtension) {
			continue
		}
		s := &Slide{
			name:         f.Name(),
			content:      data,
			presentation: p,
		}
		p.slides = append(p.slides, s)
	}
	return p, nil
}

func (p *Presentation) parseREADMEmd(data string) error {
	logger.Info("parsing README file")
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, string(t1)) && p.title == nil {
			p.title = &TElement{
				label:   tt,
				rawLine: strings.TrimPrefix(line, string(t1)),
			}
			logger.Info("title: %s", p.title.RawString())
			continue
		}
		if strings.HasPrefix(line, string(t2)) {
			if p.author == nil {
				p.author = &TElement{
					label:   ll,
					rawLine: strings.TrimPrefix(line, string(t2)),
				}
				logger.Info("author: %s", p.author.RawString())

			} else {
				p.meta = &TElement{
					label:   rr,
					rawLine: strings.TrimPrefix(line, string(t2)),
				}
				logger.Info("meta: %s", p.meta.RawString())

			}
		}
	}
	//logger.Info("Title: %s, Author: %s, Meta: %s", p.title.RawString(), p.author.RawString(), p.meta.RawString())
	return nil
}

func (p *Presentation) SetTheme(theme Theme) error {
	p.theme = theme
	return nil
}

func (p *Presentation) RenderPresentation() error {
	logger.Info("rendering %d slides", len(p.slides))
	for _, slide := range p.slides {
		pslide, err := RenderSlide(slide)
		if err != nil {
			return fmt.Errorf("unable to render slide %s: %v", slide.name, err)
		}
		p.renderedSlides = append(p.renderedSlides, pslide)
	}
	return nil
}

func (p *Presentation) DisplayPresentation() error {
	logger.Info("displaying slides")

	startingPoint := 0
rewind:
	for k := startingPoint; k < len(p.renderedSlides); k++ {
		rslide := p.renderedSlides[k]
		action, err := rslide.DisplayWithOptions(&DisplayOptions{})
		if err != nil {
			return fmt.Errorf("unable to render slide %s: %v", rslide.slide.name, err)
		}
		switch action.(type) {
		case DisplayActionPrev:
			if k > 0 {
				startingPoint = k - 1
			}
			goto rewind
		case DisplayActionStart:
			startingPoint = 0
			goto rewind
		case DisplayActionEnd:
			startingPoint = len(p.renderedSlides) - 1
			goto rewind
		}
	}
	return nil
}
