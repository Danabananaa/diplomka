// install (please make sure versions match peerDependencies)
// yarn add @nivo/core @nivo/pie
import { ResponsivePie } from '@nivo/pie'

// make sure parent container have a defined height when using
// responsive component, otherwise height will be 0 and
// no chart will be rendered.
// website examples showcase many properties,
// you'll often use just a few of them.
const theme = {
    
}

const DonutChart = ({ data /* see data tab */ }) => (
    <ResponsivePie
        data={data}
        theme={{
            fontSize: 16,
            
           }}
        margin={{ top: 40, right: 80, bottom: 80, left: 80 }}
        sortByValue={true}
        innerRadius={0.7}
        padAngle={1}
        cornerRadius={3}
        activeInnerRadiusOffset={3}
        activeOuterRadiusOffset={8}
        colors={{ scheme: 'category10' }}
        borderColor={{
            from: 'color',
            modifiers: [
                [
                    'darker',
                    0.2
                ]
            ]
        }}
        arcLinkLabelsTextColor= 'black'
        arcLinkLabelsThickness={5}
        arcLinkLabelsColor={{ from: 'color' }}
        arcLabelsSkipAngle={18}
        motionConfig="wobbly"
        legends={[
            {
                anchor: 'bottom',
                direction: 'row',
                justify: false,
                translateX: 0,
                translateY: 69,
                itemsSpacing: 0,
                itemWidth: 108,
                itemHeight: 33,
                itemTextColor: 'black',
                itemDirection: 'left-to-right',
                itemOpacity: 1,
                symbolSize: 20,
                symbolShape: 'circle',
                effects: [
                    {
                        on: 'hover',
                        style: {
                            itemTextColor: 'white'
                        }
                    }
                ]
            }
        ]}
    />
)
export default DonutChart;