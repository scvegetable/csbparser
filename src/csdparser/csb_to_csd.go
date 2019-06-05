package csdparser

import (
    "csbparser"
    "fmt"
    "math"

    "github.com/beevik/etree"
)

func CreateFromCSB(csb *csbparser.CSParseBinary) *etree.Element {
    content := etree.NewElement("Content")
    if nodeAction := csb.Action(nil); nodeAction != nil {
        content.AddChild(createAnimation(nodeAction))
    }
    animationLength := csb.AnimationListLength()
    animationList := etree.NewElement("AnimationList")
    for i := 0; i < animationLength; i++ {
        animationInfo := &csbparser.AnimationInfo{}
        if csb.AnimationList(animationInfo, i) {
            elem := createAnimationInfo(animationInfo)
            animationList.AddChild(elem)
        }
    }
    if len(animationList.Child) > 0 {
        content.AddChild(animationList)
    }
    if nodeTree := csb.NodeTree(nil); nodeTree != nil {
        content.AddChild(createObjectData(nodeTree, true))
    }
    return content
}

func createAnimation(csb *csbparser.NodeAction) *etree.Element {
    animation := etree.NewElement("Animation")
    animation.CreateAttr("Duration", intValue(int64(csb.Duration())))
    setFloatAttr(animation, "Speed", float64(csb.Speed()), 0)
    setStringAttr(animation, "ActivedAnimationName", csb.CurrentAnimationName(), "")
    timeLinesLength := csb.TimeLinesLength()
    for i := 0; i < timeLinesLength; i++ {
        timeLine := &csbparser.TimeLine{}
        if csb.TimeLines(timeLine, i) {
            animation.AddChild(createTimeLine(timeLine))
        }
    }
    return animation
}

func createTimeLine(csb *csbparser.TimeLine) *etree.Element {
    elem := etree.NewElement("Timeline")
    setIntAttr(elem, "ActionTag", int64(csb.ActionTag()), 0)
    setStringAttr(elem, "Property", csb.Property(), "")
    framesLength := csb.FramesLength()
    for i := 0; i < framesLength; i++ {
        frame := &csbparser.Frame{}
        if csb.Frames(frame, i) {
            elem.AddChild(createFrame(frame))
        }
    }
    return elem
}

func createFrame(csb *csbparser.Frame) *etree.Element {
    if pointFrame := csb.PointFrame(nil); pointFrame != nil {
        return createPointFrame(pointFrame)
    } else if scaleFrame := csb.ScaleFrame(nil); scaleFrame != nil {
        return createScaleFrame(scaleFrame)
    } else if colorFrame := csb.ColorFrame(nil); colorFrame != nil {
        return createColorFrame(colorFrame)
    } else if textureFrame := csb.TextureFrame(nil); textureFrame != nil {
        return createTextureFrame(textureFrame)
    } else if eventFrame := csb.EventFrame(nil); eventFrame != nil {
        return createEventFrame(eventFrame)
    } else if intFrame := csb.IntFrame(nil); intFrame != nil {
        return createIntFrame(intFrame)
    } else if boolFrame := csb.BoolFrame(nil); boolFrame != nil {
        return createBoolFrame(boolFrame)
    } else if innerActionFrame := csb.InnerActionFrame(nil); innerActionFrame != nil {
        return createInnerActionFrame(innerActionFrame)
    } else if blendFrame := csb.BlendFrame(nil); blendFrame != nil {
        return createBlendFrame(blendFrame)
    }
    return etree.NewElement("Frame")
}

func createPointFrame(csb *csbparser.PointFrame) *etree.Element {
    elem := etree.NewElement("PointFrame")
    elem.CreateAttr("FrameIndex", intValue(int64(csb.FrameIndex())))
    setBoolAttr(elem, "Tween", csb.Tween(), true)
    if position := csb.Position(nil); position != nil {
        elem.CreateAttr("X", floatValue(position.X()))
        elem.CreateAttr("Y", floatValue(position.Y()))
    }
    addEasingData(elem, csb.EasingData(nil))
    return elem
}

func createScaleFrame(csb *csbparser.ScaleFrame) *etree.Element {
    elem := etree.NewElement("ScaleFrame")
    elem.CreateAttr("FrameIndex", intValue(int64(csb.FrameIndex())))
    setBoolAttr(elem, "Tween", csb.Tween(), true)
    if scale := csb.Scale(nil); scale != nil {
        elem.CreateAttr("X", floatValue(scale.ScaleX()))
        elem.CreateAttr("Y", floatValue(scale.ScaleY()))
    }
    addEasingData(elem, csb.EasingData(nil))
    return elem
}

func createColorFrame(csb *csbparser.ColorFrame) *etree.Element {
    elem := etree.NewElement("ColorFrame")
    elem.CreateAttr("FrameIndex", intValue(int64(csb.FrameIndex())))
    setBoolAttr(elem, "Tween", csb.Tween(), true)
    color := etree.NewElement("Color")
    initColor(color, csb.Color(nil))
    addEasingData(elem, csb.EasingData(nil))
    return elem
}

func createTextureFrame(csb *csbparser.TextureFrame) *etree.Element {
    elem := etree.NewElement("TextureFrame")
    elem.CreateAttr("FrameIndex", intValue(int64(csb.FrameIndex())))
    setBoolAttr(elem, "Tween", csb.Tween(), true)
    addResourceData(elem, "TextureFile", csb.TextureFile(nil))
    return elem
}

func createEventFrame(csb *csbparser.EventFrame) *etree.Element {
    elem := etree.NewElement("EventFrame")
    elem.CreateAttr("FrameIndex", intValue(int64(csb.FrameIndex())))
    setBoolAttr(elem, "Tween", csb.Tween(), true)
    elem.CreateAttr("Value", csb.Value())
    addEasingData(elem, csb.EasingData(nil))
    return elem
}

func createIntFrame(csb *csbparser.IntFrame) *etree.Element {
    elem := etree.NewElement("IntFrame")
    elem.CreateAttr("FrameIndex", intValue(int64(csb.FrameIndex())))
    setBoolAttr(elem, "Tween", csb.Tween(), true)
    elem.CreateAttr("Value", intValue(int64(csb.Value())))
    addEasingData(elem, csb.EasingData(nil))
    return elem
}

func createBoolFrame(csb *csbparser.BoolFrame) *etree.Element {
    elem := etree.NewElement("BoolFrame")
    elem.CreateAttr("FrameIndex", intValue(int64(csb.FrameIndex())))
    setBoolAttr(elem, "Tween", csb.Tween(), true)
    elem.CreateAttr("Value", boolValue(csb.Value()))
    addEasingData(elem, csb.EasingData(nil))
    return elem
}

func createInnerActionFrame(csb *csbparser.InnerActionFrame) *etree.Element {
    elem := etree.NewElement("InnerActionFrame")
    elem.CreateAttr("FrameIndex", intValue(int64(csb.FrameIndex())))
    setBoolAttr(elem, "Tween", csb.Tween(), true)
    switch csb.InnerActionType() {
    case 0:
        elem.CreateAttr("InnerActionType", "LoopAction")
    case 1:
        elem.CreateAttr("InnerActionType", "NoLoopAction")
    case 2:
        elem.CreateAttr("InnerActionType", "SingleFrame")
    }
    elem.CreateAttr("CurrentAniamtionName", csb.CurrentAniamtionName())
    elem.CreateAttr("SingleFrameIndex", intValue(int64(csb.SingleFrameIndex())))
    addEasingData(elem, csb.EasingData(nil))
    return elem
}

func createBlendFrame(csb *csbparser.BlendFrame) *etree.Element {
    elem := etree.NewElement("BlendFuncFrame")
    elem.CreateAttr("FrameIndex", intValue(int64(csb.FrameIndex())))
    setBoolAttr(elem, "Tween", csb.Tween(), true)
    if blendFunc := csb.BlendFunc(nil); blendFunc != nil {
        elem.CreateAttr("Src", intValue(int64(blendFunc.Src())))
        elem.CreateAttr("Dst", intValue(int64(blendFunc.Dst())))
    }
    addEasingData(elem, csb.EasingData(nil))
    return elem
}

func addEasingData(elem *etree.Element, csb *csbparser.EasingData) {
    if csb == nil {
        return
    }
    easingData := etree.NewElement("EasingData")
    easingData.CreateAttr("Type", intValue(int64(csb.EaseType())))
    pointsLength := csb.PointsLength()
    points := etree.NewElement("Points")
    for i := 0; i < pointsLength; i++ {
        point := &csbparser.Position{}
        if csb.Points(point, i) {
            point_ := etree.NewElement("Position")
            point_.CreateAttr("X", floatValue(point.X()))
            point_.CreateAttr("Y", floatValue(point.Y()))
            points.AddChild(point_)
        }
    }
    if len(points.Child) > 0 {
        elem.AddChild(points)
    }
    elem.AddChild(easingData)
}

func createAnimationInfo(csb *csbparser.AnimationInfo) *etree.Element {
    elem := etree.NewElement("AnimationInfo")
    elem.CreateAttr("Name", csb.Name())
    elem.CreateAttr("StartIndex", intValue(int64(csb.StartIndex())))
    elem.CreateAttr("EndIndex", intValue(int64(csb.EndIndex())))

    renderColor := etree.NewElement("RenderColor")
    renderColor.CreateAttr("A", "255")
    renderColor.CreateAttr("R", "255")
    renderColor.CreateAttr("G", "255")
    renderColor.CreateAttr("B", "255")
    elem.AddChild(renderColor)
    return elem
}

func createObjectData(csb *csbparser.NodeTree, isRoot bool) *etree.Element {
    objectData := etree.NewElement("ObjectData")
    objectData.CreateAttr("Name", "")
    setStringAttr(objectData, "CustomClassName", csb.CustomClassName(), "")
    if isRoot {
    } else {
        objectData.Tag = "AbstractNodeData"
    }

    objectData.AddChild(etree.NewElement("Size"))

    childrenLength := csb.ChildrenLength()
    children := etree.NewElement("Children")
    for i := 0; i < childrenLength; i++ {
        nodeTree := &csbparser.NodeTree{}
        if csb.Children(nodeTree, i) {
            children.AddChild(createObjectData(nodeTree, false))
        }
    }
    if len(children.Child) > 0 {
        objectData.AddChild(children)
    }

    if options := csb.Options(nil); options != nil {
        initOptions(objectData, options, csb.Classname())
    }
    ctype := csb.Classname() + "ObjectData"
    if isRoot {
        ctype = "Game" + ctype
    }
    if isRoot {
        objectData.RemoveAttr("IconVisible")
        for len(objectData.Child) > 2 {
            objectData.RemoveChildAt(2)
        }
    }
    objectData.CreateAttr("ctype", ctype)

    return objectData
}

func initOptions(elem *etree.Element, csb *csbparser.Options, classname string) {
    tab := csb.Data(nil).Table()
    switch classname {
    case "Node":
        options := &csbparser.WidgetOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options)
    case "Widget":
        options := &csbparser.WidgetOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options)
    case "SingleNode":
        options := &csbparser.SingleNodeOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.NodeOptions(nil))
        initSingleNodeOptions(elem, options)
    case "Sprite":
        options := &csbparser.SpriteOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.NodeOptions(nil))
        initSpriteOptions(elem, options)
    case "Particle":
        options := &csbparser.ParticleSystemOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.NodeOptions(nil))
        initParticleSystemOptions(elem, options)
    case "GameMap":
        options := &csbparser.GameMapOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.NodeOptions(nil))
        initGameMapOptions(elem, options)
    case "TextButton":
        options := &csbparser.ButtonOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initButtonOptions(elem, options)
    case "Button":
        options := &csbparser.ButtonOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initButtonOptions(elem, options)
    case "CheckBox":
        options := &csbparser.CheckBoxOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initCheckBoxOptions(elem, options)
    case "ImageView":
        options := &csbparser.ImageViewOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initImageViewOptions(elem, options)
    case "LabelAtlas":
        options := &csbparser.TextAtlasOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initTextAtlasOptions(elem, options)
    case "TextAtlas":
        options := &csbparser.TextAtlasOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initTextAtlasOptions(elem, options)
    case "LabelBMFont":
        options := &csbparser.TextBMFontOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initTextBMFontOptions(elem, options)
    case "TextBMFont":
        options := &csbparser.TextBMFontOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initTextBMFontOptions(elem, options)
    case "TextArea":
        options := &csbparser.TextOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initTextOptions(elem, options)
    case "Label":
        options := &csbparser.TextOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initTextOptions(elem, options)
    case "Text":
        options := &csbparser.TextOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initTextOptions(elem, options)
    case "TextField":
        options := &csbparser.TextFieldOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initTextFieldOptions(elem, options)
    case "LoadingBar":
        options := &csbparser.LoadingBarOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initLoadingBarOptions(elem, options)
    case "Slider":
        options := &csbparser.SliderOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initSliderOptions(elem, options)
    case "Panel":
        options := &csbparser.PanelOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initPanelOptions(elem, options)
    case "ScrollView":
        options := &csbparser.ScrollViewOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initScrollViewOptions(elem, options)
    case "PageView":
        options := &csbparser.PageViewOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initPageViewOptions(elem, options)
    case "ListView":
        options := &csbparser.ListViewOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.WidgetOptions(nil))
        initListViewOptions(elem, options)
    case "ProjectNode":
        options := &csbparser.ProjectNodeOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.NodeOptions(nil))
        initProjectNodeOptions(elem, options)
    case "Component":
        options := &csbparser.ComponentOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.NodeOptions(nil))
        initComponentOptions(elem, options)
    case "ComAudio":
        options := &csbparser.ComAudioOptions{}
        options.Init(tab.Bytes, tab.Pos)
        initWidgetOptions(elem, options.NodeOptions(nil))
        initComAudioOptions(elem, options)
    default:
        fmt.Println("initOptions: unknown class name", classname)
    }
}

func initWidgetOptions(elem *etree.Element, csb *csbparser.WidgetOptions) {
    elem.CreateAttr("Name", csb.Name())
    setIntAttr(elem, "ActionTag", int64(csb.ActionTag()), 0)
    setBoolAttr(elem, "VisibleForFrame", csb.Visible(), true)
    setByteAttr(elem, "Alpha", csb.Alpha(), 255)
    setStringAttr(elem, "CallBackType", csb.CallBackType(), "")
    setStringAttr(elem, "CallBackName", csb.CallBackName(), "")
    setStringAttr(elem, "UserData", csb.CustomProperty(), "")
    setIntAttr(elem, "ZOrder", int64(csb.ZOrder()), 0)
    setIntAttr(elem, "Tag", int64(csb.Tag()), 0)
    setStringAttr(elem, "FrameEvent", csb.FrameEvent(), "")
    if rotationSkew := csb.RotationSkew(nil); rotationSkew != nil {
        setFloatAttr(elem, "RotationSkewX", float64(rotationSkew.RotationSkewX()), 0)
        setFloatAttr(elem, "RotationSkewY", float64(rotationSkew.RotationSkewY()), 0)
    }
    elem.CreateAttr("IconVisible", "False")
    if layoutComponent := csb.LayoutComponent(nil); layoutComponent != nil {
        setBoolAttr(elem, "PositionPercentXEnabled", layoutComponent.PositionXPercentEnabled(), false)
        setBoolAttr(elem, "PositionPercentYEnabled", layoutComponent.PositionYPercentEnabled(), false)
        setBoolAttr(elem, "PercentWidthEnable", layoutComponent.SizeXPercentEnable(), false)
        setBoolAttr(elem, "PercentHeightEnable", layoutComponent.SizeYPercentEnable(), false)
        setBoolAttr(elem, "PercentWidthEnabled", layoutComponent.SizeXPercentEnable(), false)
        setBoolAttr(elem, "PercentHeightEnabled", layoutComponent.SizeYPercentEnable(), false)
        setStringAttr(elem, "HorizontalEdge", layoutComponent.HorizontalEdge(), "")
        setStringAttr(elem, "VerticalEdge", layoutComponent.VerticalEdge(), "")
        setFloatAttr(elem, "LeftMargin", float64(layoutComponent.LeftMargin()), 0)
        setFloatAttr(elem, "RightMargin", float64(layoutComponent.RightMargin()), 0)
        setFloatAttr(elem, "TopMargin", float64(layoutComponent.TopMargin()), 0)
        setFloatAttr(elem, "BottomMargin", float64(layoutComponent.BottomMargin()), 0)
    }
    setBoolAttr(elem, "TouchEnable", csb.TouchEnabled(), false)
    if layoutComponent := csb.LayoutComponent(nil); layoutComponent != nil {
        setBoolAttr(elem, "StretchWidthEnable", layoutComponent.StretchHorizontalEnabled(), false)
        setBoolAttr(elem, "StretchHeightEnable", layoutComponent.StretchVerticalEnabled(), false)
    }
    setBoolAttr(elem, "FlipX", csb.FlipX(), false)
    setBoolAttr(elem, "FlipY", csb.FlipY(), false)

    size := elem.FindElement("Size")
    if size_ := csb.Size(nil); size_ != nil {
        size.CreateAttr("X", floatValue(size_.Width()))
        size.CreateAttr("Y", floatValue(size_.Height()))
    }
    anchorPoint := etree.NewElement("AnchorPoint")
    if anchorPoint_ := csb.AnchorPoint(nil); anchorPoint_ != nil {
        setFloatAttr(anchorPoint, "ScaleX", float64(anchorPoint_.ScaleX()), 0)
        setFloatAttr(anchorPoint, "ScaleY", float64(anchorPoint_.ScaleY()), 0)
    }
    elem.AddChild(anchorPoint)
    position := etree.NewElement("Position")
    initPosition(position, csb.Position(nil))
    elem.AddChild(position)
    scale := etree.NewElement("Scale")
    if scale_ := csb.Scale(nil); scale_ != nil {
        scale.CreateAttr("ScaleX", floatValue(scale_.ScaleX()))
        scale.CreateAttr("ScaleY", floatValue(scale_.ScaleY()))
    }
    elem.AddChild(scale)
    ccolor := etree.NewElement("CColor")
    initColor(ccolor, csb.Color(nil))
    elem.AddChild(ccolor)
    prePosition := etree.NewElement("PrePosition")
    preSize := etree.NewElement("PreSize")
    if layoutComponent := csb.LayoutComponent(nil); layoutComponent != nil {
        setFloatAttr(prePosition, "X", float64(layoutComponent.PositionXPercent()), 0)
        setFloatAttr(prePosition, "Y", float64(layoutComponent.PositionYPercent()), 0)
        preSize.CreateAttr("X", floatValue(layoutComponent.SizeXPercent()))
        preSize.CreateAttr("Y", floatValue(layoutComponent.SizeYPercent()))
    }
    elem.AddChild(prePosition)
    elem.AddChild(preSize)
}

func initSingleNodeOptions(elem *etree.Element, csb *csbparser.SingleNodeOptions) {
}

func initSpriteOptions(elem *etree.Element, csb *csbparser.SpriteOptions) {
    addResourceData(elem, "FileData", csb.FileNameData(nil))
    blendFunc := etree.NewElement("BlendFunc")
    initBlendFunc(blendFunc, csb.BlendFunc(nil))
    elem.AddChild(blendFunc)
}

func initParticleSystemOptions(elem *etree.Element, csb *csbparser.ParticleSystemOptions) {
    addResourceData(elem, "FileData", csb.FileNameData(nil))
    blendFunc := etree.NewElement("BlendFunc")
    initBlendFunc(blendFunc, csb.BlendFunc(nil))
    elem.AddChild(blendFunc)
}

func initGameMapOptions(elem *etree.Element, csb *csbparser.GameMapOptions) {
    addResourceData(elem, "FileData", csb.FileNameData(nil))
}

func initButtonOptions(elem *etree.Element, csb *csbparser.ButtonOptions) {
    setIntAttr(elem, "FontSize", int64(csb.FontSize()), 0)
    setStringAttr(elem, "FontName", csb.FontName(), "")
    setStringAttr(elem, "ButtonText", csb.Text(), "")
    setBoolAttr(elem, "Scale9Enable", csb.Scale9Enabled(), false)
    initCapInsets(elem, csb.CapInsets(nil))
    setBoolAttr(elem, "DisplayState", csb.Displaystate(), true)
    setIntAttr(elem, "OutlineSize", int64(csb.OutlineSize()), 1)
    setBoolAttr(elem, "OutlineEnabled", csb.OutlineEnabled(), false)
    elem.CreateAttr("ShadowOffsetX", floatValue(csb.ShadowOffsetX()))
    elem.CreateAttr("ShadowOffsetY", floatValue(csb.ShadowOffsetY()))
    setBoolAttr(elem, "ShadowEnabled", csb.ShadowEnabled(), false)
    setIntAttr(elem, "ShadowBlurRadius", int64(csb.ShadowBlurRadius()), 0)

    addResourceData(elem, "FontResource", csb.FontResource(nil))
    textColor := etree.NewElement("TextColor")
    initColor(textColor, csb.TextColor(nil))
    elem.AddChild(textColor)
    addResourceData(elem, "DisabledFileData", csb.DisabledData(nil))
    addResourceData(elem, "PressedFileData", csb.PressedData(nil))
    addResourceData(elem, "NormalFileData", csb.NormalData(nil))
    outlineColor := etree.NewElement("OutlineColor")
    initColor(outlineColor, csb.OutlineColor(nil))
    elem.AddChild(outlineColor)
    shadowColor := etree.NewElement("ShadowColor")
    initColor(shadowColor, csb.ShadowColor(nil))
    elem.AddChild(shadowColor)
}

func initCheckBoxOptions(elem *etree.Element, csb *csbparser.CheckBoxOptions) {
    setBoolAttr(elem, "CheckedState", csb.SelectedState(), false)
    setBoolAttr(elem, "DisplayState", csb.Displaystate(), true)

    addResourceData(elem, "NormalBackFileData", csb.BackGroundBoxData(nil))
    addResourceData(elem, "PressedBackFileData", csb.BackGroundBoxSelectedData(nil))
    addResourceData(elem, "DisableBackFileData", csb.BackGroundBoxDisabledData(nil))
    addResourceData(elem, "NodeNormalFileData", csb.FrontCrossData(nil))
    addResourceData(elem, "NodeDisableFileData", csb.FrontCrossDisabledData(nil))
}

func initImageViewOptions(elem *etree.Element, csb *csbparser.ImageViewOptions) {
    setBoolAttr(elem, "Scale9Enable", csb.Scale9Enabled(), false)
    initCapInsets(elem, csb.CapInsets(nil))

    addResourceData(elem, "FileData", csb.FileNameData(nil))
}

func initTextAtlasOptions(elem *etree.Element, csb *csbparser.TextAtlasOptions) {
    elem.CreateAttr("CharWidth", intValue(int64(csb.ItemWidth())))
    elem.CreateAttr("CharHeight", intValue(int64(csb.ItemHeight())))
    elem.CreateAttr("LabelText", csb.StringValue())
    elem.CreateAttr("StartChar", csb.StartCharMap())

    addResourceData(elem, "LabelAtlasFileImage_CNB", csb.CharMapFileData(nil))
}

func initTextBMFontOptions(elem *etree.Element, csb *csbparser.TextBMFontOptions) {
    setStringAttr(elem, "LabelText", csb.Text(), "")

    addResourceData(elem, "LabelBMFontFile_CNB", csb.FileNameData(nil))
}

func initTextOptions(elem *etree.Element, csb *csbparser.TextOptions) {
    setBoolAttr(elem, "TouchScaleChangeAble", csb.TouchScaleEnable(), false)
    setBoolAttr(elem, "IsCustomSize", csb.IsCustomSize(), false)
    setIntAttr(elem, "FontSize", int64(csb.FontSize()), 0)
    setStringAttr(elem, "FontName", csb.FontName(), "")
    elem.CreateAttr("LabelText", csb.Text())
    setIntAttr(elem, "AreaWidth", int64(csb.AreaWidth()), 0)
    setIntAttr(elem, "AreaHeight", int64(csb.AreaHeight()), 0)
    switch csb.HAlignment() {
    case 1:
        elem.CreateAttr("HorizontalAlignmentType", "HT_Center")
    case 2:
        elem.CreateAttr("HorizontalAlignmentType", "HT_Right")
    }
    switch csb.VAlignment() {
    case 1:
        elem.CreateAttr("VerticalAlignmentType", "VT_Center")
    case 2:
        elem.CreateAttr("VerticalAlignmentType", "VT_Bottom")
    }
    setIntAttr(elem, "OutlineSize", int64(csb.OutlineSize()), 1)
    setBoolAttr(elem, "OutlineEnabled", csb.OutlineEnabled(), false)
    elem.CreateAttr("ShadowOffsetX", floatValue(csb.ShadowOffsetX()))
    elem.CreateAttr("ShadowOffsetY", floatValue(csb.ShadowOffsetY()))
    setBoolAttr(elem, "ShadowEnabled", csb.ShadowEnabled(), false)
    setIntAttr(elem, "ShadowBlurRadius", int64(csb.ShadowBlurRadius()), 0)

    addResourceData(elem, "FontResource", csb.FontResource(nil))
    outlineColor := etree.NewElement("OutlineColor")
    initColor(outlineColor, csb.OutlineColor(nil))
    elem.AddChild(outlineColor)
    shadowColor := etree.NewElement("ShadowColor")
    initColor(shadowColor, csb.ShadowColor(nil))
    elem.AddChild(shadowColor)
}

func initTextFieldOptions(elem *etree.Element, csb *csbparser.TextFieldOptions) {
    setIntAttr(elem, "FontSize", int64(csb.FontSize()), 0)
    setBoolAttr(elem, "IsCustomSize", csb.IsCustomSize(), false)
    setStringAttr(elem, "FontName", csb.FontName(), "")
    elem.CreateAttr("LabelText", csb.Text())
    elem.CreateAttr("PlaceHolderText", csb.PlaceHolder())
    setBoolAttr(elem, "MaxLengthEnable", csb.MaxLengthEnabled(), false)
    setIntAttr(elem, "MaxLengthText", int64(csb.MaxLength()), 0)
    setBoolAttr(elem, "PasswordEnable", csb.PasswordEnabled(), false)
    setStringAttr(elem, "PasswordStyleText", csb.PasswordStyleText(), "*")

    addResourceData(elem, "FontResource", csb.FontResource(nil))
}

func initLoadingBarOptions(elem *etree.Element, csb *csbparser.LoadingBarOptions) {
    setIntAttr(elem, "ProgressInfo", int64(csb.Percent()), 80)
    switch csb.Direction() {
    case 1:
        elem.CreateAttr("ProgressType", "Right_To_Left")
    }

    addResourceData(elem, "ImageFileData", csb.TextureData(nil))
}

func initSliderOptions(elem *etree.Element, csb *csbparser.SliderOptions) {
    setIntAttr(elem, "PercentInfo", int64(csb.Percent()), 0)
    setBoolAttr(elem, "DisplayState", csb.Displaystate(), true)

    addResourceData(elem, "BackGroundData", csb.BarFileNameData(nil))
    addResourceData(elem, "ProgressBarData", csb.ProgressBarData(nil))
    addResourceData(elem, "BallNormalData", csb.BallNormalData(nil))
    addResourceData(elem, "BallPressedData", csb.BallPressedData(nil))
    addResourceData(elem, "BallDisabledData", csb.BallDisabledData(nil))
}

func initPanelOptions(elem *etree.Element, csb *csbparser.PanelOptions) {
    elem.CreateAttr("ClipAble", boolValue(csb.ClipEnabled()))
    setUintAttr(elem, "BackColorAlpha", uint64(csb.BgColorOpacity()), 255)
    setIntAttr(elem, "ComboBoxIndex", int64(csb.ColorType()), 0)
    initColorAngle(elem, csb.ColorVector(nil))
    setBoolAttr(elem, "Scale9Enable", csb.BackGroundScale9Enabled(), false)
    initCapInsets(elem, csb.CapInsets(nil))

    addResourceData(elem, "FileData", csb.BackGroundImageData(nil))
    singleColor := etree.NewElement("SingleColor")
    initColor(singleColor, csb.BgColor(nil))
    elem.AddChild(singleColor)
    firstColor := etree.NewElement("FirstColor")
    initColor(firstColor, csb.BgStartColor(nil))
    elem.AddChild(firstColor)
    endColor := etree.NewElement("EndColor")
    initColor(endColor, csb.BgEndColor(nil))
    elem.AddChild(endColor)
    colorVector := etree.NewElement("ColorVector")
    initColorVector(colorVector, csb.ColorVector(nil))
    elem.AddChild(colorVector)
}

func initScrollViewOptions(elem *etree.Element, csb *csbparser.ScrollViewOptions) {
    elem.CreateAttr("ClipAble", boolValue(csb.ClipEnabled()))
    setUintAttr(elem, "BackColorAlpha", uint64(csb.BgColorOpacity()), 255)
    setIntAttr(elem, "ComboBoxIndex", int64(csb.ColorType()), 0)
    initColorAngle(elem, csb.ColorVector(nil))
    setBoolAttr(elem, "Scale9Enable", csb.BackGroundScale9Enabled(), false)
    initCapInsets(elem, csb.CapInsets(nil))
    setBoolAttr(elem, "IsBounceEnabled", csb.BounceEnabled(), false)
    setBoolAttr(elem, "BarEnabled", csb.ScrollbarEnabeld(), true)
    setBoolAttr(elem, "BarAutoHide", csb.ScrollbarAutoHide(), true)
    setFloatAttr(elem, "BarAutoHideTime", float64(csb.ScrollbarAutoHideTime()), 0.2)
    switch csb.Direction() {
    case 1:
        elem.CreateAttr("ScrollDirectionType", "Vertical")
    case 2:
        elem.CreateAttr("ScrollDirectionType", "Horizontal")
    case 3:
        elem.CreateAttr("ScrollDirectionType", "Vertical_Horizontal")
    }

    addResourceData(elem, "FileData", csb.BackGroundImageData(nil))
    singleColor := etree.NewElement("SingleColor")
    initColor(singleColor, csb.BgColor(nil))
    elem.AddChild(singleColor)
    firstColor := etree.NewElement("FirstColor")
    initColor(firstColor, csb.BgStartColor(nil))
    elem.AddChild(firstColor)
    endColor := etree.NewElement("EndColor")
    initColor(endColor, csb.BgEndColor(nil))
    elem.AddChild(endColor)
    colorVector := etree.NewElement("ColorVector")
    initColorVector(colorVector, csb.ColorVector(nil))
    elem.AddChild(colorVector)
    innerNodeSize := etree.NewElement("InnerNodeSize")
    if innerSize := csb.InnerSize(nil); innerSize != nil {
        innerNodeSize.CreateAttr("Width", intValue(int64(innerSize.Width())))
        innerNodeSize.CreateAttr("Height", intValue(int64(innerSize.Height())))
    }
    elem.AddChild(innerNodeSize)
}

func initPageViewOptions(elem *etree.Element, csb *csbparser.PageViewOptions) {
    elem.CreateAttr("ClipAble", boolValue(csb.ClipEnabled()))
    setUintAttr(elem, "BackColorAlpha", uint64(csb.BgColorOpacity()), 255)
    setIntAttr(elem, "ComboBoxIndex", int64(csb.ColorType()), 0)
    initColorAngle(elem, csb.ColorVector(nil))
    setBoolAttr(elem, "Scale9Enable", csb.BackGroundScale9Enabled(), false)
    initCapInsets(elem, csb.CapInsets(nil))
    setStringAttr(elem, "ScrollDirectionType", "0", "")

    addResourceData(elem, "FileData", csb.BackGroundImageData(nil))
    singleColor := etree.NewElement("SingleColor")
    initColor(singleColor, csb.BgColor(nil))
    elem.AddChild(singleColor)
    firstColor := etree.NewElement("FirstColor")
    initColor(firstColor, csb.BgStartColor(nil))
    elem.AddChild(firstColor)
    endColor := etree.NewElement("EndColor")
    initColor(endColor, csb.BgEndColor(nil))
    elem.AddChild(endColor)
    colorVector := etree.NewElement("ColorVector")
    initColorVector(colorVector, csb.ColorVector(nil))
    elem.AddChild(colorVector)
}

func initListViewOptions(elem *etree.Element, csb *csbparser.ListViewOptions) {
    elem.CreateAttr("ClipAble", boolValue(csb.ClipEnabled()))
    setUintAttr(elem, "BackColorAlpha", uint64(csb.BgColorOpacity()), 255)
    setIntAttr(elem, "ComboBoxIndex", int64(csb.ColorType()), 0)
    initColorAngle(elem, csb.ColorVector(nil))
    setBoolAttr(elem, "Scale9Enable", csb.BackGroundScale9Enabled(), false)
    initCapInsets(elem, csb.CapInsets(nil))
    setBoolAttr(elem, "IsBounceEnabled", csb.BounceEnabled(), false)
    setStringAttr(elem, "ScrollDirectionType", "0", "")
    setIntAttr(elem, "ItemMargin", int64(csb.ItemMargin()), 0)
    setStringAttr(elem, "DirectionType", csb.DirectionType(), "")
    setStringAttr(elem, "HorizontalType", csb.HorizontalType(), "")
    setStringAttr(elem, "VerticalType", csb.VerticalType(), "")

    addResourceData(elem, "FileData", csb.BackGroundImageData(nil))
    singleColor := etree.NewElement("SingleColor")
    initColor(singleColor, csb.BgColor(nil))
    elem.AddChild(singleColor)
    firstColor := etree.NewElement("FirstColor")
    initColor(firstColor, csb.BgStartColor(nil))
    elem.AddChild(firstColor)
    endColor := etree.NewElement("EndColor")
    initColor(endColor, csb.BgEndColor(nil))
    elem.AddChild(endColor)
    colorVector := etree.NewElement("ColorVector")
    initColorVector(colorVector, csb.ColorVector(nil))
    elem.AddChild(colorVector)
}

func initProjectNodeOptions(elem *etree.Element, csb *csbparser.ProjectNodeOptions) {
    elem.CreateAttr("InnerActionSpeed", floatValue(csb.InnerActionSpeed()))

    fileData := etree.NewElement("FileData")
    fileData.CreateAttr("Type", "Normal")
    fileData.CreateAttr("Path", csb.FileName())
    fileData.CreateAttr("Plist", "")
    elem.AddChild(fileData)
}

func initComponentOptions(elem *etree.Element, csb *csbparser.ComponentOptions) {
}

func initComAudioOptions(elem *etree.Element, csb *csbparser.ComAudioOptions) {
    setStringAttr(elem, "Name", csb.Name(), "")
    setIntAttr(elem, "Volume", int64(csb.Volume()), 0)
    setBoolAttr(elem, "Loop", csb.Loop(), false)

    addResourceData(elem, "FileData", csb.FileNameData(nil))
}

func initPosition(elem *etree.Element, csb *csbparser.Position) {
    if csb == nil {
        return
    }
    setFloatAttr(elem, "X", float64(csb.X()), 0)
    setFloatAttr(elem, "Y", float64(csb.Y()), 0)
}

func initColor(elem *etree.Element, csb *csbparser.Color) {
    if csb == nil {
        elem.CreateAttr("A", "0")
        elem.CreateAttr("R", "0")
        elem.CreateAttr("G", "0")
        elem.CreateAttr("B", "0")
    } else {
        elem.CreateAttr("A", byteValue(csb.A()))
        elem.CreateAttr("R", byteValue(csb.R()))
        elem.CreateAttr("G", byteValue(csb.G()))
        elem.CreateAttr("B", byteValue(csb.B()))
    }
}

func initColorAngle(elem *etree.Element, csb *csbparser.ColorVector) {
    if csb == nil {
        return
    }
    x, y := float64(csb.VectorX()), float64(csb.VectorY())
    angle := math.Atan2(y, x) / math.Pi * 180
    setFloatAttr(elem, "ColorAngle", angle, 0)
}

func initColorVector(elem *etree.Element, csb *csbparser.ColorVector) {
    if csb == nil {
        return
    }
    setFloatAttr(elem, "ScaleX", float64(csb.VectorX()), 0)
    setFloatAttr(elem, "ScaleY", float64(csb.VectorY()), 0)
}

func initCapInsets(elem *etree.Element, csb *csbparser.CapInsets) {
    if csb == nil {
        return
    }
    setIntAttr(elem, "LeftEage", int64(math.Abs(float64(csb.X()))), 0)
    setIntAttr(elem, "RightEage", int64(math.Abs(float64(csb.X()))), 0)
    setIntAttr(elem, "TopEage", int64(math.Abs(float64(csb.Y()))), 0)
    setIntAttr(elem, "BottomEage", int64(math.Abs(float64(csb.Y()))), 0)
    setIntAttr(elem, "Scale9OriginX", int64(csb.X()), 0)
    setIntAttr(elem, "Scale9OriginY", int64(csb.Y()), 0)
    setIntAttr(elem, "Scale9Width", int64(csb.Width()), 0)
    setIntAttr(elem, "Scale9Height", int64(csb.Height()), 0)
}

func initBlendFunc(elem *etree.Element, csb *csbparser.BlendFunc) {
    if csb == nil {
        return
    }
    setIntAttr(elem, "Src", int64(csb.Src()), 0)
    setIntAttr(elem, "Dst", int64(csb.Dst()), 0)
}

func addResourceData(elem *etree.Element, name string, csb *csbparser.ResourceData) {
    if csb == nil ||csb.Path() == "" {
        return
    }
    resourceData := etree.NewElement(name)
    if csb.ResourceType() == 0 {
        if csb.Path()[:8] == "Default/" {
            resourceData.CreateAttr("Type", "Default")
        } else {
            resourceData.CreateAttr("Type", "Normal")
        }
    } else {
        resourceData.CreateAttr("Type", "PlistSubImage")
    }
    resourceData.CreateAttr("Path", csb.Path())
    resourceData.CreateAttr("Plist", csb.PlistFile())
    elem.AddChild(resourceData)
}

func boolValue(value bool) string {
    if !value {
        return "False"
    }
    return "True"
}

func byteValue(value byte) string {
    return fmt.Sprintf("%d", value)
}

func intValue(value int64) string {
    return fmt.Sprintf("%d", value)
}

func floatValue(value float32, args ...interface{}) string {
    return fmt.Sprintf("%.4f", value)
}

func setByteAttr(elem *etree.Element, name string, value, dflt byte) *etree.Attr {
    if value == dflt {
        return elem.RemoveAttr(name)
    }
    return elem.CreateAttr(name, fmt.Sprintf("%d", value))
}

func setIntAttr(elem *etree.Element, name string, value, dflt int64) *etree.Attr {
    if value == dflt {
        return elem.RemoveAttr(name)
    }
    return elem.CreateAttr(name, fmt.Sprintf("%d", value))
}

func setUintAttr(elem *etree.Element, name string, value, dflt uint64) *etree.Attr {
    if value == dflt {
        return elem.RemoveAttr(name)
    }
    return elem.CreateAttr(name, fmt.Sprintf("%d", value))
}

func setFloatAttr(elem *etree.Element, name string, value, dflt float64) *etree.Attr {
    if value == dflt {
        return elem.RemoveAttr(name)
    }
    return elem.CreateAttr(name, fmt.Sprintf("%.4f", value))
}

func setBoolAttr(elem *etree.Element, name string, value, dflt bool) *etree.Attr {
    if value == dflt {
        return elem.RemoveAttr(name)
    }
    if value {
        return elem.CreateAttr(name, "True")
    } else {
        return elem.CreateAttr(name, "False")
    }
}

func setStringAttr(elem *etree.Element, name string, value, dflt string) *etree.Attr {
    if value == dflt {
        return elem.RemoveAttr(name)
    }
    return elem.CreateAttr(name, value)
}
