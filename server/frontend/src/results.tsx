import { Person, Voice } from './types'

export const Results = ({
  seenIn,
  person,
}: {
  seenIn: Voice[]
  person: Person
}) => {
  const namePieces = person.name.split(' ')

  return (
    <>
      <div className="divider"></div>
      <div style={{ display: 'flex', justifyContent: 'flex-end' }}>
        <div
          style={{
            display: 'flex',
            flexDirection: 'column',
            justifyContent: 'flex-end',
          }}
        >
          {namePieces.map((np) => (
            <h2
              style={{
                width: '100%',
                fontSize: '2rem',
                textAlign: 'center',
              }}
            >
              {np}
            </h2>
          ))}
        </div>
        <img
          src={person.images.jpg.image_url}
          style={{
            maxWidth: '50%',
            display: 'block',
            marginRight: 8,
          }}
        />
      </div>

      {seenIn.map((si) => (
        <>
          <div className="divider"></div>
          <div style={{ display: 'flex' }}>
            <img
              src={si.character.images.jpg.image_url}
              style={{
                marginLeft: 8,
                width: '50vw',
                display: 'inline-block',
              }}
            />

            <div
              style={{
                width: '50vw',
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'flex-end',
                marginLeft: 8,
              }}
            >
              <a href={si.anime.url} target="_blank">
                <h3
                  style={{
                    width: '100%',
                    fontSize: '1.5rem',
                    textAlign: 'left',
                    fontStyle: 'italic',
                    marginBottom: 8,
                  }}
                >
                  {si.anime.title}
                </h3>
              </a>

              <h2 style={{ fontSize: '2rem' }}>
                {si.character.name}
              </h2>
            </div>
          </div>
        </>
      ))}
    </>
  )
}
