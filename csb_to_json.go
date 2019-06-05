package main

import (
    "csbparser"
    "fmt"
    flatbuffers"github.com/google/flatbuffers/go"
)

type Decoder interface {
    GetData() interface{}
}

func ConvertOptions(options *csbparser.Options, classname string) Decoder {
    tab := options.Data(nil).Table()
    switch classname {
    case "Node":
        return &WidgetOptions{_tab: tab}
    case "Widget":
        return &WidgetOptions{_tab: tab}
    case "SingleNode":
        return &SingleNodeOptions{_tab: tab}
    case "Sprite":
        return &SpriteOptions{_tab: tab}
    case "Particle":
        return &ParticleSystemOptions{_tab: tab}
    case "GameMap":
        return &GameMapOptions{_tab: tab}
    case "TextButton":
        return &ButtonOptions{_tab: tab}
    case "Button":
        return &ButtonOptions{_tab: tab}
    case "CheckBox":
        return &CheckBoxOptions{_tab: tab}
    case "ImageView":
        return &ImageViewOptions{_tab: tab}
    case "LabelAtlas":
        return &TextAtlasOptions{_tab: tab}
    case "TextAtlas":
        return &TextAtlasOptions{_tab: tab}
    case "LabelBMFont":
        return &TextBMFontOptions{_tab: tab}
    case "TextBMFont":
        return &TextBMFontOptions{_tab: tab}
    case "TextArea":
        return &TextOptions{_tab: tab}
    case "Label":
        return &TextOptions{_tab: tab}
    case "Text":
        return &TextOptions{_tab: tab}
    case "TextField":
        return &TextFieldOptions{_tab: tab}
    case "LoadingBar":
        return &LoadingBarOptions{_tab: tab}
    case "Slider":
        return &SliderOptions{_tab: tab}
    case "Panel":
        return &PanelOptions{_tab: tab}
    case "ScrollView":
        return &ScrollViewOptions{_tab: tab}
    case "PageView":
        return &PageViewOptions{_tab: tab}
    case "ListView":
        return &ListViewOptions{_tab: tab}
    case "ProjectNode":
        return &ProjectNodeOptions{_tab: tab}
    case "Component":
        return &ComponentOptions{_tab: tab}
    case "ComAudio":
        return &ComAudioOptions{_tab:tab}
    }
    fmt.Println("unknown class name", classname)
    return nil
}
type CSParseBinary struct { _tab flatbuffers.Table }
func (this *CSParseBinary)GetData() interface{} {
    table := &csbparser.CSParseBinary{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["version"] = table.Version()
    texturesLen := table.TexturesLength()
    textures := []string{}
    for i := 0; i < texturesLen; i++ {
        textures = append(textures, table.Textures(i))
    }
    ret["textures"] = textures
    texturePngsLen := table.TexturePngsLength()
    texturePngs := []string{}
    for i := 0; i < texturePngsLen; i++ {
        texturePngs = append(texturePngs, table.TexturePngs(i))
    }
    ret["texturePngs"] = texturePngs
    if nodeTree := table.NodeTree(nil); nodeTree != nil {
        ret["nodeTree"] = (&NodeTree{_tab: nodeTree.Table()}).GetData()
    }
    if action := table.Action(nil); action != nil {
        ret["action"] = (&NodeAction{_tab: action.Table()}).GetData()
    }
    animationListLen := table.AnimationListLength()
    animationList := []interface{}{}
    for i := 0; i < animationListLen; i++ {
        value := &csbparser.AnimationInfo{}
        if table.AnimationList(value, i) {
            animationList = append(animationList, (&AnimationInfo{_tab: value.Table()}).GetData())
        }
    }
    ret["animationList"] = animationList
    return ret
}
type NodeTree struct { _tab flatbuffers.Table }
func (this *NodeTree)GetData() interface{} {
    table := &csbparser.NodeTree{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["classname"] = table.Classname()
    childrenLen := table.ChildrenLength()
    children := []interface{}{}
    for i := 0; i < childrenLen; i++ {
        value := &csbparser.NodeTree{}
        if table.Children(value, i) {
            children = append(children, (&NodeTree{_tab: value.Table()}).GetData())
        }
    }
    ret["children"] = children
    if options := table.Options(nil); options != nil {
        ret["options"] = ConvertOptions(options, table.Classname()).GetData()
    }
    ret["customClassName"] = table.CustomClassName()
    return ret
}
type Options struct { _tab flatbuffers.Table }
func (this *Options)GetData() interface{} {
    table := &csbparser.Options{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if data := table.Data(nil); data != nil {
        ret["data"] = (&WidgetOptions{_tab: data.Table()}).GetData()
    }
    return ret
}
type WidgetOptions struct { _tab flatbuffers.Table }
func (this *WidgetOptions)GetData() interface{} {
    table := &csbparser.WidgetOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["name"] = table.Name()
    ret["actionTag"] = table.ActionTag()
    if rotationSkew := table.RotationSkew(nil); rotationSkew != nil {
        ret["rotationSkew"] = (&RotationSkew{_tab: rotationSkew.Table()}).GetData()
    }
    ret["zOrder"] = table.ZOrder()
    ret["visible"] = table.Visible()
    ret["alpha"] = table.Alpha()
    ret["tag"] = table.Tag()
    if position := table.Position(nil); position != nil {
        ret["position"] = (&Position{_tab: position.Table()}).GetData()
    }
    if scale := table.Scale(nil); scale != nil {
        ret["scale"] = (&Scale{_tab: scale.Table()}).GetData()
    }
    if anchorPoint := table.AnchorPoint(nil); anchorPoint != nil {
        ret["anchorPoint"] = (&AnchorPoint{_tab: anchorPoint.Table()}).GetData()
    }
    if color := table.Color(nil); color != nil {
        ret["color"] = (&Color{_tab: color.Table()}).GetData()
    }
    if size := table.Size(nil); size != nil {
        ret["size"] = (&FlatSize{_tab: size.Table()}).GetData()
    }
    ret["flipX"] = table.FlipX()
    ret["flipY"] = table.FlipY()
    ret["ignoreSize"] = table.IgnoreSize()
    ret["touchEnabled"] = table.TouchEnabled()
    ret["frameEvent"] = table.FrameEvent()
    ret["customProperty"] = table.CustomProperty()
    ret["callBackType"] = table.CallBackType()
    ret["callBackName"] = table.CallBackName()
    if layoutComponent := table.LayoutComponent(nil); layoutComponent != nil {
        ret["layoutComponent"] = (&LayoutComponentTable{_tab: layoutComponent.Table()}).GetData()
    }
    return ret
}
type LayoutComponentTable struct { _tab flatbuffers.Table }
func (this *LayoutComponentTable)GetData() interface{} {
    table := &csbparser.LayoutComponentTable{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["positionXPercentEnabled"] = table.PositionXPercentEnabled()
    ret["positionYPercentEnabled"] = table.PositionYPercentEnabled()
    ret["positionXPercent"] = table.PositionXPercent()
    ret["positionYPercent"] = table.PositionYPercent()
    ret["sizeXPercentEnable"] = table.SizeXPercentEnable()
    ret["sizeYPercentEnable"] = table.SizeYPercentEnable()
    ret["sizeXPercent"] = table.SizeXPercent()
    ret["sizeYPercent"] = table.SizeYPercent()
    ret["stretchHorizontalEnabled"] = table.StretchHorizontalEnabled()
    ret["stretchVerticalEnabled"] = table.StretchVerticalEnabled()
    ret["horizontalEdge"] = table.HorizontalEdge()
    ret["verticalEdge"] = table.VerticalEdge()
    ret["leftMargin"] = table.LeftMargin()
    ret["rightMargin"] = table.RightMargin()
    ret["topMargin"] = table.TopMargin()
    ret["bottomMargin"] = table.BottomMargin()
    return ret
}
type SingleNodeOptions struct { _tab flatbuffers.Table }
func (this *SingleNodeOptions)GetData() interface{} {
    table := &csbparser.SingleNodeOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if nodeOptions := table.NodeOptions(nil); nodeOptions != nil {
        ret["nodeOptions"] = (&WidgetOptions{_tab: nodeOptions.Table()}).GetData()
    }
    return ret
}
type SpriteOptions struct { _tab flatbuffers.Table }
func (this *SpriteOptions)GetData() interface{} {
    table := &csbparser.SpriteOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if nodeOptions := table.NodeOptions(nil); nodeOptions != nil {
        ret["nodeOptions"] = (&WidgetOptions{_tab: nodeOptions.Table()}).GetData()
    }
    if fileNameData := table.FileNameData(nil); fileNameData != nil {
        ret["fileNameData"] = (&ResourceData{_tab: fileNameData.Table()}).GetData()
    }
    if blendFunc := table.BlendFunc(nil); blendFunc != nil {
        ret["blendFunc"] = (&BlendFunc{_tab: blendFunc.Table()}).GetData()
    }
    return ret
}
type ParticleSystemOptions struct { _tab flatbuffers.Table }
func (this *ParticleSystemOptions)GetData() interface{} {
    table := &csbparser.ParticleSystemOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if nodeOptions := table.NodeOptions(nil); nodeOptions != nil {
        ret["nodeOptions"] = (&WidgetOptions{_tab: nodeOptions.Table()}).GetData()
    }
    if fileNameData := table.FileNameData(nil); fileNameData != nil {
        ret["fileNameData"] = (&ResourceData{_tab: fileNameData.Table()}).GetData()
    }
    if blendFunc := table.BlendFunc(nil); blendFunc != nil {
        ret["blendFunc"] = (&BlendFunc{_tab: blendFunc.Table()}).GetData()
    }
    return ret
}
type GameMapOptions struct { _tab flatbuffers.Table }
func (this *GameMapOptions)GetData() interface{} {
    table := &csbparser.GameMapOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if nodeOptions := table.NodeOptions(nil); nodeOptions != nil {
        ret["nodeOptions"] = (&WidgetOptions{_tab: nodeOptions.Table()}).GetData()
    }
    if fileNameData := table.FileNameData(nil); fileNameData != nil {
        ret["fileNameData"] = (&ResourceData{_tab: fileNameData.Table()}).GetData()
    }
    return ret
}
type ButtonOptions struct { _tab flatbuffers.Table }
func (this *ButtonOptions)GetData() interface{} {
    table := &csbparser.ButtonOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if widgetOptions := table.WidgetOptions(nil); widgetOptions != nil {
        ret["widgetOptions"] = (&WidgetOptions{_tab: widgetOptions.Table()}).GetData()
    }
    if normalData := table.NormalData(nil); normalData != nil {
        ret["normalData"] = (&ResourceData{_tab: normalData.Table()}).GetData()
    }
    if pressedData := table.PressedData(nil); pressedData != nil {
        ret["pressedData"] = (&ResourceData{_tab: pressedData.Table()}).GetData()
    }
    if disabledData := table.DisabledData(nil); disabledData != nil {
        ret["disabledData"] = (&ResourceData{_tab: disabledData.Table()}).GetData()
    }
    if fontResource := table.FontResource(nil); fontResource != nil {
        ret["fontResource"] = (&ResourceData{_tab: fontResource.Table()}).GetData()
    }
    ret["text"] = table.Text()
    ret["fontName"] = table.FontName()
    ret["fontSize"] = table.FontSize()
    if textColor := table.TextColor(nil); textColor != nil {
        ret["textColor"] = (&Color{_tab: textColor.Table()}).GetData()
    }
    if capInsets := table.CapInsets(nil); capInsets != nil {
        ret["capInsets"] = (&CapInsets{_tab: capInsets.Table()}).GetData()
    }
    if scale9Size := table.Scale9Size(nil); scale9Size != nil {
        ret["scale9Size"] = (&FlatSize{_tab: scale9Size.Table()}).GetData()
    }
    ret["scale9Enabled"] = table.Scale9Enabled()
    ret["displaystate"] = table.Displaystate()
    ret["outlineEnabled"] = table.OutlineEnabled()
    if outlineColor := table.OutlineColor(nil); outlineColor != nil {
        ret["outlineColor"] = (&Color{_tab: outlineColor.Table()}).GetData()
    }
    ret["outlineSize"] = table.OutlineSize()
    ret["shadowEnabled"] = table.ShadowEnabled()
    if shadowColor := table.ShadowColor(nil); shadowColor != nil {
        ret["shadowColor"] = (&Color{_tab: shadowColor.Table()}).GetData()
    }
    ret["shadowOffsetX"] = table.ShadowOffsetX()
    ret["shadowOffsetY"] = table.ShadowOffsetY()
    ret["shadowBlurRadius"] = table.ShadowBlurRadius()
    return ret
}
type CheckBoxOptions struct { _tab flatbuffers.Table }
func (this *CheckBoxOptions)GetData() interface{} {
    table := &csbparser.CheckBoxOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if widgetOptions := table.WidgetOptions(nil); widgetOptions != nil {
        ret["widgetOptions"] = (&WidgetOptions{_tab: widgetOptions.Table()}).GetData()
    }
    if backGroundBoxData := table.BackGroundBoxData(nil); backGroundBoxData != nil {
        ret["backGroundBoxData"] = (&ResourceData{_tab: backGroundBoxData.Table()}).GetData()
    }
    if backGroundBoxSelectedData := table.BackGroundBoxSelectedData(nil); backGroundBoxSelectedData != nil {
        ret["backGroundBoxSelectedData"] = (&ResourceData{_tab: backGroundBoxSelectedData.Table()}).GetData()
    }
    if frontCrossData := table.FrontCrossData(nil); frontCrossData != nil {
        ret["frontCrossData"] = (&ResourceData{_tab: frontCrossData.Table()}).GetData()
    }
    if backGroundBoxDisabledData := table.BackGroundBoxDisabledData(nil); backGroundBoxDisabledData != nil {
        ret["backGroundBoxDisabledData"] = (&ResourceData{_tab: backGroundBoxDisabledData.Table()}).GetData()
    }
    if frontCrossDisabledData := table.FrontCrossDisabledData(nil); frontCrossDisabledData != nil {
        ret["frontCrossDisabledData"] = (&ResourceData{_tab: frontCrossDisabledData.Table()}).GetData()
    }
    ret["selectedState"] = table.SelectedState()
    ret["displaystate"] = table.Displaystate()
    return ret
}
type ImageViewOptions struct { _tab flatbuffers.Table }
func (this *ImageViewOptions)GetData() interface{} {
    table := &csbparser.ImageViewOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if widgetOptions := table.WidgetOptions(nil); widgetOptions != nil {
        ret["widgetOptions"] = (&WidgetOptions{_tab: widgetOptions.Table()}).GetData()
    }
    if fileNameData := table.FileNameData(nil); fileNameData != nil {
        ret["fileNameData"] = (&ResourceData{_tab: fileNameData.Table()}).GetData()
    }
    if capInsets := table.CapInsets(nil); capInsets != nil {
        ret["capInsets"] = (&CapInsets{_tab: capInsets.Table()}).GetData()
    }
    if scale9Size := table.Scale9Size(nil); scale9Size != nil {
        ret["scale9Size"] = (&FlatSize{_tab: scale9Size.Table()}).GetData()
    }
    ret["scale9Enabled"] = table.Scale9Enabled()
    return ret
}
type TextAtlasOptions struct { _tab flatbuffers.Table }
func (this *TextAtlasOptions)GetData() interface{} {
    table := &csbparser.TextAtlasOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if widgetOptions := table.WidgetOptions(nil); widgetOptions != nil {
        ret["widgetOptions"] = (&WidgetOptions{_tab: widgetOptions.Table()}).GetData()
    }
    if charMapFileData := table.CharMapFileData(nil); charMapFileData != nil {
        ret["charMapFileData"] = (&ResourceData{_tab: charMapFileData.Table()}).GetData()
    }
    ret["stringValue"] = table.StringValue()
    ret["startCharMap"] = table.StartCharMap()
    ret["itemWidth"] = table.ItemWidth()
    ret["itemHeight"] = table.ItemHeight()
    return ret
}
type TextBMFontOptions struct { _tab flatbuffers.Table }
func (this *TextBMFontOptions)GetData() interface{} {
    table := &csbparser.TextBMFontOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if widgetOptions := table.WidgetOptions(nil); widgetOptions != nil {
        ret["widgetOptions"] = (&WidgetOptions{_tab: widgetOptions.Table()}).GetData()
    }
    if fileNameData := table.FileNameData(nil); fileNameData != nil {
        ret["fileNameData"] = (&ResourceData{_tab: fileNameData.Table()}).GetData()
    }
    ret["text"] = table.Text()
    return ret
}
type TextOptions struct { _tab flatbuffers.Table }
func (this *TextOptions)GetData() interface{} {
    table := &csbparser.TextOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if widgetOptions := table.WidgetOptions(nil); widgetOptions != nil {
        ret["widgetOptions"] = (&WidgetOptions{_tab: widgetOptions.Table()}).GetData()
    }
    if fontResource := table.FontResource(nil); fontResource != nil {
        ret["fontResource"] = (&ResourceData{_tab: fontResource.Table()}).GetData()
    }
    ret["fontName"] = table.FontName()
    ret["fontSize"] = table.FontSize()
    ret["text"] = table.Text()
    ret["areaWidth"] = table.AreaWidth()
    ret["areaHeight"] = table.AreaHeight()
    ret["hAlignment"] = table.HAlignment()
    ret["vAlignment"] = table.VAlignment()
    ret["touchScaleEnable"] = table.TouchScaleEnable()
    ret["isCustomSize"] = table.IsCustomSize()
    ret["outlineEnabled"] = table.OutlineEnabled()
    if outlineColor := table.OutlineColor(nil); outlineColor != nil {
        ret["outlineColor"] = (&Color{_tab: outlineColor.Table()}).GetData()
    }
    ret["outlineSize"] = table.OutlineSize()
    ret["shadowEnabled"] = table.ShadowEnabled()
    if shadowColor := table.ShadowColor(nil); shadowColor != nil {
        ret["shadowColor"] = (&Color{_tab: shadowColor.Table()}).GetData()
    }
    ret["shadowOffsetX"] = table.ShadowOffsetX()
    ret["shadowOffsetY"] = table.ShadowOffsetY()
    ret["shadowBlurRadius"] = table.ShadowBlurRadius()
    return ret
}
type TextFieldOptions struct { _tab flatbuffers.Table }
func (this *TextFieldOptions)GetData() interface{} {
    table := &csbparser.TextFieldOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if widgetOptions := table.WidgetOptions(nil); widgetOptions != nil {
        ret["widgetOptions"] = (&WidgetOptions{_tab: widgetOptions.Table()}).GetData()
    }
    if fontResource := table.FontResource(nil); fontResource != nil {
        ret["fontResource"] = (&ResourceData{_tab: fontResource.Table()}).GetData()
    }
    ret["fontName"] = table.FontName()
    ret["fontSize"] = table.FontSize()
    ret["text"] = table.Text()
    ret["placeHolder"] = table.PlaceHolder()
    ret["passwordEnabled"] = table.PasswordEnabled()
    ret["passwordStyleText"] = table.PasswordStyleText()
    ret["maxLengthEnabled"] = table.MaxLengthEnabled()
    ret["maxLength"] = table.MaxLength()
    ret["areaWidth"] = table.AreaWidth()
    ret["areaHeight"] = table.AreaHeight()
    ret["isCustomSize"] = table.IsCustomSize()
    return ret
}
type LoadingBarOptions struct { _tab flatbuffers.Table }
func (this *LoadingBarOptions)GetData() interface{} {
    table := &csbparser.LoadingBarOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if widgetOptions := table.WidgetOptions(nil); widgetOptions != nil {
        ret["widgetOptions"] = (&WidgetOptions{_tab: widgetOptions.Table()}).GetData()
    }
    if textureData := table.TextureData(nil); textureData != nil {
        ret["textureData"] = (&ResourceData{_tab: textureData.Table()}).GetData()
    }
    ret["percent"] = table.Percent()
    ret["direction"] = table.Direction()
    return ret
}
type SliderOptions struct { _tab flatbuffers.Table }
func (this *SliderOptions)GetData() interface{} {
    table := &csbparser.SliderOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if widgetOptions := table.WidgetOptions(nil); widgetOptions != nil {
        ret["widgetOptions"] = (&WidgetOptions{_tab: widgetOptions.Table()}).GetData()
    }
    if barFileNameData := table.BarFileNameData(nil); barFileNameData != nil {
        ret["barFileNameData"] = (&ResourceData{_tab: barFileNameData.Table()}).GetData()
    }
    if ballNormalData := table.BallNormalData(nil); ballNormalData != nil {
        ret["ballNormalData"] = (&ResourceData{_tab: ballNormalData.Table()}).GetData()
    }
    if ballPressedData := table.BallPressedData(nil); ballPressedData != nil {
        ret["ballPressedData"] = (&ResourceData{_tab: ballPressedData.Table()}).GetData()
    }
    if ballDisabledData := table.BallDisabledData(nil); ballDisabledData != nil {
        ret["ballDisabledData"] = (&ResourceData{_tab: ballDisabledData.Table()}).GetData()
    }
    if progressBarData := table.ProgressBarData(nil); progressBarData != nil {
        ret["progressBarData"] = (&ResourceData{_tab: progressBarData.Table()}).GetData()
    }
    ret["percent"] = table.Percent()
    ret["displaystate"] = table.Displaystate()
    return ret
}
type PanelOptions struct { _tab flatbuffers.Table }
func (this *PanelOptions)GetData() interface{} {
    table := &csbparser.PanelOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if widgetOptions := table.WidgetOptions(nil); widgetOptions != nil {
        ret["widgetOptions"] = (&WidgetOptions{_tab: widgetOptions.Table()}).GetData()
    }
    if backGroundImageData := table.BackGroundImageData(nil); backGroundImageData != nil {
        ret["backGroundImageData"] = (&ResourceData{_tab: backGroundImageData.Table()}).GetData()
    }
    ret["clipEnabled"] = table.ClipEnabled()
    if bgColor := table.BgColor(nil); bgColor != nil {
        ret["bgColor"] = (&Color{_tab: bgColor.Table()}).GetData()
    }
    if bgStartColor := table.BgStartColor(nil); bgStartColor != nil {
        ret["bgStartColor"] = (&Color{_tab: bgStartColor.Table()}).GetData()
    }
    if bgEndColor := table.BgEndColor(nil); bgEndColor != nil {
        ret["bgEndColor"] = (&Color{_tab: bgEndColor.Table()}).GetData()
    }
    ret["colorType"] = table.ColorType()
    ret["bgColorOpacity"] = table.BgColorOpacity()
    if colorVector := table.ColorVector(nil); colorVector != nil {
        ret["colorVector"] = (&ColorVector{_tab: colorVector.Table()}).GetData()
    }
    if capInsets := table.CapInsets(nil); capInsets != nil {
        ret["capInsets"] = (&CapInsets{_tab: capInsets.Table()}).GetData()
    }
    if scale9Size := table.Scale9Size(nil); scale9Size != nil {
        ret["scale9Size"] = (&FlatSize{_tab: scale9Size.Table()}).GetData()
    }
    ret["backGroundScale9Enabled"] = table.BackGroundScale9Enabled()
    return ret
}
type ScrollViewOptions struct { _tab flatbuffers.Table }
func (this *ScrollViewOptions)GetData() interface{} {
    table := &csbparser.ScrollViewOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if widgetOptions := table.WidgetOptions(nil); widgetOptions != nil {
        ret["widgetOptions"] = (&WidgetOptions{_tab: widgetOptions.Table()}).GetData()
    }
    if backGroundImageData := table.BackGroundImageData(nil); backGroundImageData != nil {
        ret["backGroundImageData"] = (&ResourceData{_tab: backGroundImageData.Table()}).GetData()
    }
    ret["clipEnabled"] = table.ClipEnabled()
    if bgColor := table.BgColor(nil); bgColor != nil {
        ret["bgColor"] = (&Color{_tab: bgColor.Table()}).GetData()
    }
    if bgStartColor := table.BgStartColor(nil); bgStartColor != nil {
        ret["bgStartColor"] = (&Color{_tab: bgStartColor.Table()}).GetData()
    }
    if bgEndColor := table.BgEndColor(nil); bgEndColor != nil {
        ret["bgEndColor"] = (&Color{_tab: bgEndColor.Table()}).GetData()
    }
    ret["colorType"] = table.ColorType()
    ret["bgColorOpacity"] = table.BgColorOpacity()
    if colorVector := table.ColorVector(nil); colorVector != nil {
        ret["colorVector"] = (&ColorVector{_tab: colorVector.Table()}).GetData()
    }
    if capInsets := table.CapInsets(nil); capInsets != nil {
        ret["capInsets"] = (&CapInsets{_tab: capInsets.Table()}).GetData()
    }
    if scale9Size := table.Scale9Size(nil); scale9Size != nil {
        ret["scale9Size"] = (&FlatSize{_tab: scale9Size.Table()}).GetData()
    }
    ret["backGroundScale9Enabled"] = table.BackGroundScale9Enabled()
    if innerSize := table.InnerSize(nil); innerSize != nil {
        ret["innerSize"] = (&FlatSize{_tab: innerSize.Table()}).GetData()
    }
    ret["direction"] = table.Direction()
    ret["bounceEnabled"] = table.BounceEnabled()
    ret["scrollbarEnabeld"] = table.ScrollbarEnabeld()
    ret["scrollbarAutoHide"] = table.ScrollbarAutoHide()
    ret["scrollbarAutoHideTime"] = table.ScrollbarAutoHideTime()
    return ret
}
type PageViewOptions struct { _tab flatbuffers.Table }
func (this *PageViewOptions)GetData() interface{} {
    table := &csbparser.PageViewOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if widgetOptions := table.WidgetOptions(nil); widgetOptions != nil {
        ret["widgetOptions"] = (&WidgetOptions{_tab: widgetOptions.Table()}).GetData()
    }
    if backGroundImageData := table.BackGroundImageData(nil); backGroundImageData != nil {
        ret["backGroundImageData"] = (&ResourceData{_tab: backGroundImageData.Table()}).GetData()
    }
    ret["clipEnabled"] = table.ClipEnabled()
    if bgColor := table.BgColor(nil); bgColor != nil {
        ret["bgColor"] = (&Color{_tab: bgColor.Table()}).GetData()
    }
    if bgStartColor := table.BgStartColor(nil); bgStartColor != nil {
        ret["bgStartColor"] = (&Color{_tab: bgStartColor.Table()}).GetData()
    }
    if bgEndColor := table.BgEndColor(nil); bgEndColor != nil {
        ret["bgEndColor"] = (&Color{_tab: bgEndColor.Table()}).GetData()
    }
    ret["colorType"] = table.ColorType()
    ret["bgColorOpacity"] = table.BgColorOpacity()
    if colorVector := table.ColorVector(nil); colorVector != nil {
        ret["colorVector"] = (&ColorVector{_tab: colorVector.Table()}).GetData()
    }
    if capInsets := table.CapInsets(nil); capInsets != nil {
        ret["capInsets"] = (&CapInsets{_tab: capInsets.Table()}).GetData()
    }
    if scale9Size := table.Scale9Size(nil); scale9Size != nil {
        ret["scale9Size"] = (&FlatSize{_tab: scale9Size.Table()}).GetData()
    }
    ret["backGroundScale9Enabled"] = table.BackGroundScale9Enabled()
    return ret
}
type ListViewOptions struct { _tab flatbuffers.Table }
func (this *ListViewOptions)GetData() interface{} {
    table := &csbparser.ListViewOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if widgetOptions := table.WidgetOptions(nil); widgetOptions != nil {
        ret["widgetOptions"] = (&WidgetOptions{_tab: widgetOptions.Table()}).GetData()
    }
    if backGroundImageData := table.BackGroundImageData(nil); backGroundImageData != nil {
        ret["backGroundImageData"] = (&ResourceData{_tab: backGroundImageData.Table()}).GetData()
    }
    ret["clipEnabled"] = table.ClipEnabled()
    if bgColor := table.BgColor(nil); bgColor != nil {
        ret["bgColor"] = (&Color{_tab: bgColor.Table()}).GetData()
    }
    if bgStartColor := table.BgStartColor(nil); bgStartColor != nil {
        ret["bgStartColor"] = (&Color{_tab: bgStartColor.Table()}).GetData()
    }
    if bgEndColor := table.BgEndColor(nil); bgEndColor != nil {
        ret["bgEndColor"] = (&Color{_tab: bgEndColor.Table()}).GetData()
    }
    ret["colorType"] = table.ColorType()
    ret["bgColorOpacity"] = table.BgColorOpacity()
    if colorVector := table.ColorVector(nil); colorVector != nil {
        ret["colorVector"] = (&ColorVector{_tab: colorVector.Table()}).GetData()
    }
    if capInsets := table.CapInsets(nil); capInsets != nil {
        ret["capInsets"] = (&CapInsets{_tab: capInsets.Table()}).GetData()
    }
    if scale9Size := table.Scale9Size(nil); scale9Size != nil {
        ret["scale9Size"] = (&FlatSize{_tab: scale9Size.Table()}).GetData()
    }
    ret["backGroundScale9Enabled"] = table.BackGroundScale9Enabled()
    if innerSize := table.InnerSize(nil); innerSize != nil {
        ret["innerSize"] = (&FlatSize{_tab: innerSize.Table()}).GetData()
    }
    ret["direction"] = table.Direction()
    ret["bounceEnabled"] = table.BounceEnabled()
    ret["itemMargin"] = table.ItemMargin()
    ret["directionType"] = table.DirectionType()
    ret["horizontalType"] = table.HorizontalType()
    ret["verticalType"] = table.VerticalType()
    return ret
}
type ProjectNodeOptions struct { _tab flatbuffers.Table }
func (this *ProjectNodeOptions)GetData() interface{} {
    table := &csbparser.ProjectNodeOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if nodeOptions := table.NodeOptions(nil); nodeOptions != nil {
        ret["nodeOptions"] = (&WidgetOptions{_tab: nodeOptions.Table()}).GetData()
    }
    ret["fileName"] = table.FileName()
    ret["innerActionSpeed"] = table.InnerActionSpeed()
    return ret
}
type ComponentOptions struct { _tab flatbuffers.Table }
func (this *ComponentOptions)GetData() interface{} {
    table := &csbparser.ComponentOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if nodeOptions := table.NodeOptions(nil); nodeOptions != nil {
        ret["nodeOptions"] = (&WidgetOptions{_tab: nodeOptions.Table()}).GetData()
    }
    ret["comType"] = table.ComType()
    if comAudioOptions := table.ComAudioOptions(nil); comAudioOptions != nil {
        ret["comAudioOptions"] = (&ComAudioOptions{_tab: comAudioOptions.Table()}).GetData()
    }
    return ret
}
type ComAudioOptions struct { _tab flatbuffers.Table }
func (this *ComAudioOptions)GetData() interface{} {
    table := &csbparser.ComAudioOptions{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if nodeOptions := table.NodeOptions(nil); nodeOptions != nil {
        ret["nodeOptions"] = (&WidgetOptions{_tab: nodeOptions.Table()}).GetData()
    }
    ret["name"] = table.Name()
    ret["enabled"] = table.Enabled()
    ret["loop"] = table.Loop()
    ret["volume"] = table.Volume()
    if fileNameData := table.FileNameData(nil); fileNameData != nil {
        ret["fileNameData"] = (&ResourceData{_tab: fileNameData.Table()}).GetData()
    }
    return ret
}
type AnimationInfo struct { _tab flatbuffers.Table }
func (this *AnimationInfo)GetData() interface{} {
    table := &csbparser.AnimationInfo{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["name"] = table.Name()
    ret["startIndex"] = table.StartIndex()
    ret["endIndex"] = table.EndIndex()
    return ret
}
type NodeAction struct { _tab flatbuffers.Table }
func (this *NodeAction)GetData() interface{} {
    table := &csbparser.NodeAction{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["duration"] = table.Duration()
    ret["speed"] = table.Speed()
    timeLinesLen := table.TimeLinesLength()
    timeLines := []interface{}{}
    for i := 0; i < timeLinesLen; i++ {
        value := &csbparser.TimeLine{}
        if table.TimeLines(value, i) {
            timeLines = append(timeLines, (&TimeLine{_tab: value.Table()}).GetData())
        }
    }
    ret["timeLines"] = timeLines
    ret["currentAnimationName"] = table.CurrentAnimationName()
    return ret
}
type TimeLine struct { _tab flatbuffers.Table }
func (this *TimeLine)GetData() interface{} {
    table := &csbparser.TimeLine{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["property"] = table.Property()
    ret["actionTag"] = table.ActionTag()
    framesLen := table.FramesLength()
    frames := []interface{}{}
    for i := 0; i < framesLen; i++ {
        value := &csbparser.Frame{}
        if table.Frames(value, i) {
            frames = append(frames, (&Frame{_tab: value.Table()}).GetData())
        }
    }
    ret["frames"] = frames
    return ret
}
type Frame struct { _tab flatbuffers.Table }
func (this *Frame)GetData() interface{} {
    table := &csbparser.Frame{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    if pointFrame := table.PointFrame(nil); pointFrame != nil {
        ret["pointFrame"] = (&PointFrame{_tab: pointFrame.Table()}).GetData()
    }
    if scaleFrame := table.ScaleFrame(nil); scaleFrame != nil {
        ret["scaleFrame"] = (&ScaleFrame{_tab: scaleFrame.Table()}).GetData()
    }
    if colorFrame := table.ColorFrame(nil); colorFrame != nil {
        ret["colorFrame"] = (&ColorFrame{_tab: colorFrame.Table()}).GetData()
    }
    if textureFrame := table.TextureFrame(nil); textureFrame != nil {
        ret["textureFrame"] = (&TextureFrame{_tab: textureFrame.Table()}).GetData()
    }
    if eventFrame := table.EventFrame(nil); eventFrame != nil {
        ret["eventFrame"] = (&EventFrame{_tab: eventFrame.Table()}).GetData()
    }
    if intFrame := table.IntFrame(nil); intFrame != nil {
        ret["intFrame"] = (&IntFrame{_tab: intFrame.Table()}).GetData()
    }
    if boolFrame := table.BoolFrame(nil); boolFrame != nil {
        ret["boolFrame"] = (&BoolFrame{_tab: boolFrame.Table()}).GetData()
    }
    if innerActionFrame := table.InnerActionFrame(nil); innerActionFrame != nil {
        ret["innerActionFrame"] = (&InnerActionFrame{_tab: innerActionFrame.Table()}).GetData()
    }
    if blendFrame := table.BlendFrame(nil); blendFrame != nil {
        ret["blendFrame"] = (&BlendFrame{_tab: blendFrame.Table()}).GetData()
    }
    return ret
}
type PointFrame struct { _tab flatbuffers.Table }
func (this *PointFrame)GetData() interface{} {
    table := &csbparser.PointFrame{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["frameIndex"] = table.FrameIndex()
    ret["tween"] = table.Tween()
    if position := table.Position(nil); position != nil {
        ret["position"] = (&Position{_tab: position.Table()}).GetData()
    }
    if easingData := table.EasingData(nil); easingData != nil {
        ret["easingData"] = (&EasingData{_tab: easingData.Table()}).GetData()
    }
    return ret
}
type ScaleFrame struct { _tab flatbuffers.Table }
func (this *ScaleFrame)GetData() interface{} {
    table := &csbparser.ScaleFrame{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["frameIndex"] = table.FrameIndex()
    ret["tween"] = table.Tween()
    if scale := table.Scale(nil); scale != nil {
        ret["scale"] = (&Scale{_tab: scale.Table()}).GetData()
    }
    if easingData := table.EasingData(nil); easingData != nil {
        ret["easingData"] = (&EasingData{_tab: easingData.Table()}).GetData()
    }
    return ret
}
type ColorFrame struct { _tab flatbuffers.Table }
func (this *ColorFrame)GetData() interface{} {
    table := &csbparser.ColorFrame{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["frameIndex"] = table.FrameIndex()
    ret["tween"] = table.Tween()
    if color := table.Color(nil); color != nil {
        ret["color"] = (&Color{_tab: color.Table()}).GetData()
    }
    if easingData := table.EasingData(nil); easingData != nil {
        ret["easingData"] = (&EasingData{_tab: easingData.Table()}).GetData()
    }
    return ret
}
type TextureFrame struct { _tab flatbuffers.Table }
func (this *TextureFrame)GetData() interface{} {
    table := &csbparser.TextureFrame{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["frameIndex"] = table.FrameIndex()
    ret["tween"] = table.Tween()
    if textureFile := table.TextureFile(nil); textureFile != nil {
        ret["textureFile"] = (&ResourceData{_tab: textureFile.Table()}).GetData()
    }
    if easingData := table.EasingData(nil); easingData != nil {
        ret["easingData"] = (&EasingData{_tab: easingData.Table()}).GetData()
    }
    return ret
}
type EventFrame struct { _tab flatbuffers.Table }
func (this *EventFrame)GetData() interface{} {
    table := &csbparser.EventFrame{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["frameIndex"] = table.FrameIndex()
    ret["tween"] = table.Tween()
    ret["value"] = table.Value()
    if easingData := table.EasingData(nil); easingData != nil {
        ret["easingData"] = (&EasingData{_tab: easingData.Table()}).GetData()
    }
    return ret
}
type IntFrame struct { _tab flatbuffers.Table }
func (this *IntFrame)GetData() interface{} {
    table := &csbparser.IntFrame{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["frameIndex"] = table.FrameIndex()
    ret["tween"] = table.Tween()
    ret["value"] = table.Value()
    if easingData := table.EasingData(nil); easingData != nil {
        ret["easingData"] = (&EasingData{_tab: easingData.Table()}).GetData()
    }
    return ret
}
type BoolFrame struct { _tab flatbuffers.Table }
func (this *BoolFrame)GetData() interface{} {
    table := &csbparser.BoolFrame{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["frameIndex"] = table.FrameIndex()
    ret["tween"] = table.Tween()
    ret["value"] = table.Value()
    if easingData := table.EasingData(nil); easingData != nil {
        ret["easingData"] = (&EasingData{_tab: easingData.Table()}).GetData()
    }
    return ret
}
type InnerActionFrame struct { _tab flatbuffers.Table }
func (this *InnerActionFrame)GetData() interface{} {
    table := &csbparser.InnerActionFrame{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["frameIndex"] = table.FrameIndex()
    ret["tween"] = table.Tween()
    ret["innerActionType"] = table.InnerActionType()
    ret["currentAniamtionName"] = table.CurrentAniamtionName()
    ret["singleFrameIndex"] = table.SingleFrameIndex()
    if easingData := table.EasingData(nil); easingData != nil {
        ret["easingData"] = (&EasingData{_tab: easingData.Table()}).GetData()
    }
    return ret
}
type EasingData struct { _tab flatbuffers.Table }
func (this *EasingData)GetData() interface{} {
    table := &csbparser.EasingData{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["easeType"] = table.EaseType()
    pointsLen := table.PointsLength()
    points := []interface{}{}
    for i := 0; i < pointsLen; i++ {
        value := &csbparser.Position{}
        if table.Points(value, i) {
            points = append(points, (&Position{_tab: value.Table()}).GetData())
        }
    }
    ret["points"] = points
    return ret
}
type RotationSkew struct { _tab flatbuffers.Table }
func (this *RotationSkew)GetData() interface{} {
    table := &csbparser.RotationSkew{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["rotationSkewX"] = table.RotationSkewX()
    ret["rotationSkewY"] = table.RotationSkewY()
    return ret
}
type Position struct { _tab flatbuffers.Table }
func (this *Position)GetData() interface{} {
    table := &csbparser.Position{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["x"] = table.X()
    ret["y"] = table.Y()
    return ret
}
type Scale struct { _tab flatbuffers.Table }
func (this *Scale)GetData() interface{} {
    table := &csbparser.Scale{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["scaleX"] = table.ScaleX()
    ret["scaleY"] = table.ScaleY()
    return ret
}
type AnchorPoint struct { _tab flatbuffers.Table }
func (this *AnchorPoint)GetData() interface{} {
    table := &csbparser.AnchorPoint{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["scaleX"] = table.ScaleX()
    ret["scaleY"] = table.ScaleY()
    return ret
}
type Color struct { _tab flatbuffers.Table }
func (this *Color)GetData() interface{} {
    table := &csbparser.Color{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["a"] = table.A()
    ret["r"] = table.R()
    ret["g"] = table.G()
    ret["b"] = table.B()
    return ret
}
type ColorVector struct { _tab flatbuffers.Table }
func (this *ColorVector)GetData() interface{} {
    table := &csbparser.ColorVector{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["vectorX"] = table.VectorX()
    ret["vectorY"] = table.VectorY()
    return ret
}
type FlatSize struct { _tab flatbuffers.Table }
func (this *FlatSize)GetData() interface{} {
    table := &csbparser.FlatSize{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["width"] = table.Width()
    ret["height"] = table.Height()
    return ret
}
type CapInsets struct { _tab flatbuffers.Table }
func (this *CapInsets)GetData() interface{} {
    table := &csbparser.CapInsets{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["x"] = table.X()
    ret["y"] = table.Y()
    ret["width"] = table.Width()
    ret["height"] = table.Height()
    return ret
}
type BlendFunc struct { _tab flatbuffers.Table }
func (this *BlendFunc)GetData() interface{} {
    table := &csbparser.BlendFunc{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["src"] = table.Src()
    ret["dst"] = table.Dst()
    return ret
}
type ResourceData struct { _tab flatbuffers.Table }
func (this *ResourceData)GetData() interface{} {
    table := &csbparser.ResourceData{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["path"] = table.Path()
    ret["plistFile"] = table.PlistFile()
    ret["resourceType"] = table.ResourceType()
    return ret
}
type BlendFrame struct { _tab flatbuffers.Table }
func (this *BlendFrame)GetData() interface{} {
    table := &csbparser.BlendFrame{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
    ret["frameIndex"] = table.FrameIndex()
    ret["tween"] = table.Tween()
    if blendFunc := table.BlendFunc(nil); blendFunc != nil {
        ret["blendFunc"] = (&BlendFunc{_tab: blendFunc.Table()}).GetData()
    }
    if easingData := table.EasingData(nil); easingData != nil {
        ret["easingData"] = (&EasingData{_tab: easingData.Table()}).GetData()
    }
    return ret
}

